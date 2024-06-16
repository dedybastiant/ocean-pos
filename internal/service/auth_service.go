package service

import (
	"context"
	"database/sql"
	"errors"
	"ocean-pos/internal/dto"
	"ocean-pos/internal/model"
	"ocean-pos/internal/repository"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, request dto.LoginRequest) (*dto.LoginResponse, error)
	Logout(ctx context.Context, accessToken string, refreshToken string) (*dto.CommonResponse, error)
}

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	RDB            *redis.Client
	ViperConfig    *viper.Viper
}

func NewAuthService(repository repository.UserRepository, DB *sql.DB, RDB *redis.Client, viperConfig *viper.Viper) AuthService {
	return &AuthServiceImpl{
		UserRepository: repository,
		DB:             DB,
		RDB:            RDB,
		ViperConfig:    viperConfig,
	}
}

func GenerateAccessToken(user *model.User, viperConfig *viper.Viper) (*string, error) {
	secretKey := viperConfig.GetString("ACCESS_TOKEN_KEY")

	var jwtKey = []byte(secretKey)

	accessTokenExpiry := time.Now().Add(15 * time.Minute).Unix()

	accessTokenClaims := &model.AcessTokenClaims{
		Sub:  int(user.Id),
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExpiry,
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	accessTokenString, err := accessToken.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return &accessTokenString, nil
}

func GenerateRefreshToken(user *model.User, viperConfig *viper.Viper) (*string, error) {
	secretKey := viperConfig.GetString("REFRESH_TOKEN_KEY")

	var jwtKey = []byte(secretKey)

	refreshTokenExpiry := time.Now().Add(7 * 24 * time.Hour).Unix()

	refreshTokenClaims := &model.RefreshTokenClaims{
		Sub: int(user.Id),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshTokenExpiry,
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	refresTokenString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return &refresTokenString, nil
}

func VerifyToken(jwtToken string, secret string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("ERROR")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return &jwt.MapClaims{}, errors.New("PARSING_ERROR")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return &jwt.MapClaims{}, errors.New("INVALID_TOKEN")
	}
	return &claims, nil
}

func BlacklistToken(ctx context.Context, rdb *redis.Client, token string, expirationTimeInMinutes int64) error {
	expirationDuration := time.Minute * time.Duration(expirationTimeInMinutes)

	err := rdb.SetNX(ctx, token, true, expirationDuration).Err()
	if err != nil {
		return errors.New("FAILED_BLACKLIST_ACCOUNT")
	}
	return nil
}

func (service *AuthServiceImpl) Login(ctx context.Context, request dto.LoginRequest) (*dto.LoginResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("EMAIL_NOT_FOUND")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, errors.New("INCORRECT_CREDENTIAL")
	}

	accessToken, err := GenerateAccessToken(user, service.ViperConfig)
	refreshToken, err := GenerateRefreshToken(user, service.ViperConfig)

	response := &dto.LoginResponse{
		Token:        *accessToken,
		RefreshToken: *refreshToken,
	}

	return response, nil
}

func (service *AuthServiceImpl) Logout(ctx context.Context, accessToken string, refreshToken string) (*dto.CommonResponse, error) {
	err := BlacklistToken(ctx, service.RDB, accessToken, 5)
	if err != nil {
		return &dto.CommonResponse{}, err
	}

	err = BlacklistToken(ctx, service.RDB, refreshToken, 120)
	if err != nil {
		return &dto.CommonResponse{}, err
	}

	response := &dto.CommonResponse{
		Code:   200,
		Status: "success logout",
	}
	return response, nil
}
