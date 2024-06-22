package dto

import (
	"ocean-pos/internal/model"
	"time"
)

type CreateCategoryRequest struct {
	BusinessId int    `json:"business_id"`
	Name       string `json:"name"`
}

type CategoryResponse struct {
	Id            int        `json:"id"`
	BusinessId    int        `json:"business_id"`
	Name          string     `json:"name"`
	DeactivatedAt *time.Time `json:"deactivated_at"`
	CreatedAt     time.Time  `json:"created_at"`
	CreatedBy     int        `json:"created_by"`
	UpdatedAt     time.Time  `json:"updated_at"`
	UpdatedBy     int        `json:"updated_by"`
}

func GenerateCategoryResponse(category *model.Category) *CategoryResponse {
	return &CategoryResponse{
		Id:            category.Id,
		BusinessId:    category.BusinessId,
		Name:          category.Name,
		DeactivatedAt: category.DeactivatedAt,
		CreatedAt:     category.CreatedAt,
		CreatedBy:     category.CreatedBy,
		UpdatedAt:     category.UpdatedAt,
		UpdatedBy:     category.UpdatedBy,
	}
}
