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

type MenuService interface {
	AddNewMenu(ctx context.Context, request dto.AddMenuRequest, userId int) (*dto.MenuResponse, error)
}

type MenuServiceImpl struct {
	MenuRepository      repository.MenuRepository
	BusinessRepository  repository.BusinessRepository
	CategoryRespository repository.CategoryRespository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewMenuService(menuRepository repository.MenuRepository, businessRepository repository.BusinessRepository, categoryRespository repository.CategoryRespository, DB *sql.DB, validate *validator.Validate) MenuService {
	return &MenuServiceImpl{
		MenuRepository:      menuRepository,
		BusinessRepository:  businessRepository,
		CategoryRespository: categoryRespository,
		DB:                  DB,
		Validate:            validate,
	}
}

func (service *MenuServiceImpl) AddNewMenu(ctx context.Context, request dto.AddMenuRequest, userId int) (*dto.MenuResponse, error) {
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

	categoryId := request.CategoryId
	businessId := request.BusinessId
	menuName := request.Name
	defaultPrice := request.DefaultPrice

	_, err = service.BusinessRepository.FindBusinessByOwner(ctx, tx, userId, businessId)
	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(("BUSINESS_NOT_FOUND"))
		}
		return nil, err
	}

	category, err := service.CategoryRespository.FindCategoryById(ctx, tx, categoryId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("CATEGORY_NOT_FOUND")
		}
		return nil, err
	}

	if category.BusinessId != businessId {
		return nil, errors.New("FORBIDDEN")
	}

	exsistMenu, err := service.MenuRepository.FindMenuByName(ctx, tx, menuName)
	if exsistMenu != nil {
		return nil, errors.New("DUPLICATE_MENU")
	}

	if err == sql.ErrNoRows {
		currentTime := time.Now()

		menuData := model.Menu{
			CategoryId:   categoryId,
			Name:         menuName,
			DefaultPrice: defaultPrice,
			CreatedAt:    currentTime,
			CreatedBy:    userId,
			UpdatedAt:    currentTime,
			UpdatedBy:    userId,
		}

		menu, err := service.MenuRepository.InsertMenu(ctx, tx, menuData)
		fmt.Println(err)
		if err != nil {
			return nil, err
		}

		tx.Commit()
		return dto.GenerateMenuResponse(menu), err
	} else {
		return nil, err
	}
}
