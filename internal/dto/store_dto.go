package dto

import (
	"ocean-pos/internal/model"
	"time"
)

type RegisterStoreRequest struct {
	BusinessId  int    `json:"business_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type RegisterStoreResponse struct {
	Id            int        `json:"id"`
	BusinessId    int        `json:"business_id"`
	Name          string     `json:"name"`
	Location      string     `json:"location"`
	Description   string     `json:"description"`
	DeactivatedAt *time.Time `json:"deactivated_at"`
	CreatedAt     time.Time  `json:"created_at"`
	CreatedBy     int        `json:"created_by"`
	UpdatedAt     time.Time  `json:"updated_at"`
	UpdatedBy     int        `json:"updated_by"`
}

func GenerateRegisterStoreResponse(store *model.Store) *RegisterStoreResponse {
	return &RegisterStoreResponse{
		Id:            store.Id,
		BusinessId:    store.BusinessId,
		Name:          store.Name,
		Location:      store.Location,
		Description:   store.Description,
		DeactivatedAt: store.DeactivatedAt,
		CreatedAt:     store.CreatedAt,
		CreatedBy:     store.CreatedBy,
		UpdatedAt:     store.UpdatedAt,
		UpdatedBy:     store.UpdatedBy,
	}
}
