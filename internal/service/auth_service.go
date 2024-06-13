package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"ocean-pos/internal/dto"
	"ocean-pos/internal/model"
	"ocean-pos/internal/repository"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, request dto.AuthRequest) (*dto.AuthResponse, error)
}

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewAuthService(repository repository.UserRepository, DB *sql.DB) AuthService {
	return &AuthServiceImpl{
		UserRepository: repository,
		DB:             DB,
	}
}

var jwtKey = []byte("$2a$12$MNLqNZYZTTnS2/dvFrmmL..W6vrKSCxNS8BQaAv/jGPg6MJUCGIDm")

func GenerateToken(user *model.User) (*dto.AuthResponse, error) {
	accessTokenExpiry := time.Now().Add(15 * time.Minute).Unix()
	refreshTokenExpiry := time.Now().Add(7 * 24 * time.Hour).Unix()

	accessTokenClaims := &model.AcessTokenClaims{
		Sub:  int(user.Id),
		Name: (user.Name),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExpiry,
		},
	}

	refreshTokenClaims := &model.RefreshTokenClaims{
		Sub: int(user.Id),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshTokenExpiry,
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	accessTokenString, err := accessToken.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	refresTokenString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	token := &dto.AuthResponse{
		Token:        accessTokenString,
		RefreshToken: refresTokenString,
	}

	return token, nil
}

func VerifyToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("ERROR")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", errors.New("PARSING_ERROR")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", errors.New("INVALID_TOKEN")
	}

	userId, ok := claims["sub"].(float64)
	if !ok {
		return "", errors.New("USER_ID_NOT_FOUND")
	}
	userIdInt := int(userId)

	return strconv.Itoa(userIdInt), nil
}

func (service *AuthServiceImpl) Login(ctx context.Context, request dto.AuthRequest) (*dto.AuthResponse, error) {
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

	fmt.Println(user.Password, request.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("INCORRECT_CREDENTIAL")
	}

	response, err := GenerateToken(user)

	return response, nil
}
