package service

import (
	"context"
	"database/sql"
	"errors"
	"ocean-pos/internal/dto"
	"ocean-pos/internal/model"
	"ocean-pos/internal/repository"
	"ocean-pos/util"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, request dto.UserRequest) (*dto.UserResponse, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewUserService(repository repository.UserRepository, DB *sql.DB) UserService {
	return &UserServiceImpl{
		UserRepository: repository,
		DB:             DB,
	}
}

func hasedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashed := string(bytes)
	return hashed, nil
}

func (service *UserServiceImpl) Register(ctx context.Context, request dto.UserRequest) (*dto.UserResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	checkEmail, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if checkEmail != nil {
		return nil, errors.New("EMAIL_ALREADY_USED")
	}

	normalizePhoneNumber := util.NormalizePhoneNumber(request.PhoneNumber)
	checkPhoneNumber, err := service.UserRepository.FindByPhoneNumber(ctx, tx, normalizePhoneNumber)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if checkPhoneNumber != nil {
		return nil, errors.New("PHONE_NUMBER_ALREADY_USED")
	}

	hashedPassword, err := hasedPassword(request.Password)
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()

	userData := model.User{
		Email:       request.Email,
		Password:    hashedPassword,
		PhoneNumber: normalizePhoneNumber,
		Name:        request.Name,
		CreatedAt:   currentTime,
		CreatedBy:   1,
		UpdatedAt:   currentTime,
		UpdatedBy:   1,
	}

	user, err := service.UserRepository.Insert(ctx, tx, userData)
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return dto.GenerateUserResponse(user), nil
}
