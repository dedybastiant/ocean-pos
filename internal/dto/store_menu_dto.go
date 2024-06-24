package dto

import (
	"ocean-pos/internal/model"
	"time"
)

type AddStoreMenuRequest struct {
	BusinessId  int  `json:"business_id"`
	CategoryId  int  `json:"category_id"`
	StoreId     int  `json:"store_id"`
	MenuId      int  `json:"menu_id"`
	StorePrice  int  `json:"store_price"`
	IsAvailable bool `json:"is_available"`
}

type StoreMenuResponse struct {
	Id            int        `json:"id"`
	StoreId       int        `json:"store_id"`
	MenuId        int        `json:"menu_id"`
	StorePrice    int        `json:"store_price"`
	IsAvailable   bool       `json:"is_available"`
	DeactivatedAt *time.Time `json:"deactivated_at"`
	CreatedAt     time.Time  `json:"created_at"`
	CreatedBy     int        `json:"created_by"`
	UpdatedAt     time.Time  `json:"updated_at"`
	UpdatedBy     int        `json:"updated_by"`
}

func GenerateStoreMenuResponse(storeMenu *model.StoreMenu) *StoreMenuResponse {
	return &StoreMenuResponse{
		Id:            storeMenu.Id,
		StoreId:       storeMenu.StoreId,
		MenuId:        storeMenu.MenuId,
		StorePrice:    storeMenu.StorePrice,
		IsAvailable:   storeMenu.IsAvailable,
		DeactivatedAt: storeMenu.DeactivatedAt,
		CreatedAt:     storeMenu.CreatedAt,
		CreatedBy:     storeMenu.CreatedBy,
		UpdatedAt:     storeMenu.UpdatedAt,
		UpdatedBy:     storeMenu.UpdatedBy,
	}
}
