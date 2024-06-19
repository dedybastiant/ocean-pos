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

type BusinessService interface {
	RegisterBusiness(ctx context.Context, request dto.BusinessRequest, userId int) (*dto.BusinessResponse, error)
	GetBusinessById(ctx context.Context, businessId int, userId int) (*dto.BusinessResponse, error)
}

type BusinessServiceImpl struct {
	BusinessRepository repository.BusinessRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewBusinessService(repository repository.BusinessRepository, DB *sql.DB, validate *validator.Validate) BusinessService {
	return &BusinessServiceImpl{
		BusinessRepository: repository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *BusinessServiceImpl) RegisterBusiness(ctx context.Context, request dto.BusinessRequest, userId int) (*dto.BusinessResponse, error) {
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

	normalizePhoneNumber := util.NormalizePhoneNumber(request.PhoneNumber)

	currentTime := time.Now()

	businessData := model.Business{
		OwnerUserId: userId,
		Email:       request.Email,
		PhoneNumber: normalizePhoneNumber,
		Name:        request.Name,
		CreatedAt:   currentTime,
		CreatedBy:   userId,
		UpdatedAt:   currentTime,
		UpdatedBy:   userId,
	}

	business, err := service.BusinessRepository.Insert(ctx, tx, businessData)
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return dto.GenerateBusinessResponse(business), nil
}

func (service *BusinessServiceImpl) GetBusinessById(ctx context.Context, businessId int, userId int) (*dto.BusinessResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	result, err := service.BusinessRepository.FindBusinessById(ctx, tx, businessId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("BUSINESS_NOT_FOUND")
		}
		return nil, err
	}

	if result.OwnerUserId != userId {
		return nil, errors.New("FORBIDDEN")
	}

	return dto.GenerateBusinessResponse(result), nil
}
