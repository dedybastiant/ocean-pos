package dto

import (
	"ocean-pos/internal/model"
	"time"
)

type AddMenuRequest struct {
	BusinessId   int    `json:"business_id"`
	CategoryId   int    `json:"category_id"`
	Name         string `json:"name"`
	DefaultPrice int    `json:"default_price"`
}

type MenuResponse struct {
	Id            int        `json:"id"`
	CategoryId    int        `json:"category_id"`
	Name          string     `json:"name"`
	DefaultPrice  int        `json:"default_price"`
	DeactivatedAt *time.Time `json:"deactivated_at"`
	CreatedAt     time.Time  `json:"created_at"`
	CreatedBy     int        `json:"created_by"`
	UpdatedAt     time.Time  `json:"updated_at"`
	UpdatedBy     int        `json:"updated_by"`
}

func GenerateMenuResponse(menu *model.Menu) *MenuResponse {
	return &MenuResponse{
		Id:            menu.Id,
		CategoryId:    menu.CategoryId,
		Name:          menu.Name,
		DefaultPrice:  menu.DefaultPrice,
		DeactivatedAt: menu.DeactivatedAt,
		CreatedAt:     menu.CreatedAt,
		CreatedBy:     menu.CreatedBy,
		UpdatedAt:     menu.UpdatedAt,
		UpdatedBy:     menu.UpdatedBy,
	}
}
