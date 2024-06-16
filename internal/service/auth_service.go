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
}

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	RDB            *redis.Client
}

func NewAuthService(repository repository.UserRepository, DB *sql.DB, RDB *redis.Client) AuthService {
	return &AuthServiceImpl{
		UserRepository: repository,
		DB:             DB,
		RDB:            RDB,
	}
}

func GenerateAccessToken(user *model.User) (*string, error) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}

	secretKey := config.GetString("ACCESS_TOKEN_KEY")

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

func GenerateRefreshToken(user *model.User) (*string, error) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}

	secretKey := config.GetString("REFRESH_TOKEN_KEY")

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

	accessToken, err := GenerateAccessToken(user)
	refreshToken, err := GenerateRefreshToken(user)

	response := &dto.LoginResponse{
		Token:        *accessToken,
		RefreshToken: *refreshToken,
	}

	return response, nil
}
