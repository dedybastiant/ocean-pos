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

type StoreMenuService interface {
	AddNewStoreMenu(ctx context.Context, request dto.AddStoreMenuRequest, userId int) (*dto.StoreMenuResponse, error)
}

type StoreMenuServiceImpl struct {
	StoreMenuRepository repository.StoreMenuRepository
	BusinessRepository  repository.BusinessRepository
	StoreRepository     repository.StoreRepository
	CategoryRepository  repository.CategoryRespository
	MenuRepository      repository.MenuRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewStoreMenuService(storeMenuRepository repository.StoreMenuRepository, businessRepository repository.BusinessRepository, storeRepository repository.StoreRepository, categoryRepository repository.CategoryRespository, menuRepository repository.MenuRepository, DB *sql.DB, validate *validator.Validate) StoreMenuService {
	return &StoreMenuServiceImpl{
		StoreMenuRepository: storeMenuRepository,
		BusinessRepository:  businessRepository,
		StoreRepository:     storeRepository,
		CategoryRepository:  categoryRepository,
		MenuRepository:      menuRepository,
		DB:                  DB,
		Validate:            validate,
	}
}

func (service *StoreMenuServiceImpl) AddNewStoreMenu(ctx context.Context, request dto.AddStoreMenuRequest, userId int) (*dto.StoreMenuResponse, error) {
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

	businessId := request.BusinessId
	categoryId := request.CategoryId
	menuId := request.MenuId
	storeId := request.StoreId
	storePrice := request.StorePrice
	isAvailable := request.IsAvailable

	business, err := service.BusinessRepository.FindBusinessByOwner(ctx, tx, userId, businessId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("BUSINESS_NOT_FOUND")
		}
		return nil, err
	}

	category, err := service.CategoryRepository.FindCategoryById(ctx, tx, categoryId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("CATEGORY_NOT_FOUND")
		}
		return nil, err
	}

	if business.Id != category.BusinessId {
		return nil, errors.New("FORBIDDEN")
	}

	store, err := service.StoreRepository.FindStoreById(ctx, tx, storeId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("STORE_NOT_FOUND")
		}
		return nil, err
	}

	if business.Id != store.BusinessId {
		return nil, errors.New("FORBIDDEN")
	}

	menu, err := service.MenuRepository.FindMenuById(ctx, tx, menuId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("MENU_NOT_FOUND")
		}
		return nil, err
	}

	if category.Id != menu.CategoryId {
		return nil, errors.New("FORBIDDEN")
	}

	currentTime := time.Now()

	storeMenuData := model.StoreMenu{
		StoreId:     storeId,
		MenuId:      menuId,
		StorePrice:  storePrice,
		IsAvailable: isAvailable,
		CreatedAt:   currentTime,
		CreatedBy:   userId,
		UpdatedAt:   currentTime,
		UpdatedBy:   userId,
	}

	existStoreMenu, err := service.StoreMenuRepository.FindStoreMenuByStoreAndMenuId(ctx, tx, storeId, menuId)
	fmt.Println(err)
	if existStoreMenu != nil {
		return nil, errors.New("DUPLICATE_STORE_MENU")
	}
	if err == sql.ErrNoRows {
		storeMenu, err := service.StoreMenuRepository.InsertStoreMenu(ctx, tx, storeMenuData)
		fmt.Println(err)
		if err != nil {
			return nil, err
		}

		tx.Commit()
		return dto.GenerateStoreMenuResponse(storeMenu), nil
	} else {
		return nil, err
	}
}
