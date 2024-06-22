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

type CategoryService interface {
	CreateCategory(ctx context.Context, request dto.CreateCategoryRequest, userId int) (*dto.CategoryResponse, error)
}

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRespository
	BusinessRepository repository.BusinessRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRespository, businessRepository repository.BusinessRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		BusinessRepository: businessRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) CreateCategory(ctx context.Context, request dto.CreateCategoryRequest, userId int) (*dto.CategoryResponse, error) {
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
			return nil, errors.New("FORBIDDEN")
		}
		return nil, err
	}

	existingCategory, err := service.CategoryRepository.FindCategoryByName(ctx, tx, request.Name)
	if existingCategory != nil {
		return nil, errors.New("DUPLICATE_CATEGORY")
	}
	fmt.Println(existingCategory, err)
	if err == sql.ErrNoRows {
		currentTime := time.Now()

		categoryData := model.Category{
			BusinessId: business.Id,
			Name:       request.Name,
			CreatedAt:  currentTime,
			CreatedBy:  userId,
			UpdatedAt:  currentTime,
			UpdatedBy:  userId,
		}

		category, err := service.CategoryRepository.InsertCategory(ctx, tx, categoryData)
		fmt.Println(category, err)
		if err != nil {
			return nil, err
		}

		tx.Commit()
		return dto.GenerateCategoryResponse(category), nil
	} else {
		return nil, err
	}
}
