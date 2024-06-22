package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"ocean-pos/internal/dto"
	"ocean-pos/internal/model"
	"ocean-pos/internal/repository"
	"ocean-pos/util"
	"time"

	"github.com/go-playground/validator/v10"
)

type StoreService interface {
	RegisterStore(ctx context.Context, request dto.RegisterStoreRequest, userId int) (*dto.RegisterStoreResponse, error)
}

type StoreServiceImpl struct {
	StoreRepository    repository.StoreRepository
	BusinessRepository repository.BusinessRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewStoreService(storeRepository repository.StoreRepository, businessRepository repository.BusinessRepository, DB *sql.DB, validate *validator.Validate) StoreService {
	return &StoreServiceImpl{
		StoreRepository:    storeRepository,
		BusinessRepository: businessRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *StoreServiceImpl) RegisterStore(ctx context.Context, request dto.RegisterStoreRequest, userId int) (*dto.RegisterStoreResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := util.FormatValidationError(validationErrors)
			return nil, fmt.Errorf("validation error: %v", formattedErrors)
		}
		return nil, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	business, err := service.BusinessRepository.FindBusinessByOwner(ctx, tx, userId, request.BusinessId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("BUSINESS_NOT_FOUND")
		}
		return nil, err
	}

	currentTime := time.Now()

	storeData := model.Store{
		BusinessId:  business.Id,
		Name:        request.Name,
		Location:    request.Location,
		Description: request.Description,
		CreatedAt:   currentTime,
		CreatedBy:   userId,
		UpdatedAt:   currentTime,
		UpdatedBy:   userId,
	}

	store, err := service.StoreRepository.InsertStore(ctx, tx, storeData)
	if err != nil {
		return nil, err
	}

	return dto.GenerateRegisterStoreResponse(store), nil
}
