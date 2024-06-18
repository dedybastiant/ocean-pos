package dto

import (
	"ocean-pos/internal/model"
	"time"
)

type BusinessRequest struct {
	Email       string `validate:"required,email,min=1" json:"email"`
	PhoneNumber string `validate:"required,min=11" json:"phone_number"`
	Name        string `validate:"required,min=1" json:"name"`
}

type BusinessResponse struct {
	Id            int        `json:"id"`
	OwnerUserId   int        `json:"owner_user_id"`
	Email         string     `json:"email"`
	PhoneNumber   string     `json:"phone_number"`
	Name          string     `json:"name"`
	VerifiedAt    *time.Time `json:"verified_at"`
	DeactivatedAt *time.Time `json:"deactivated_at"`
	CreatedAt     time.Time  `json:"created_at"`
	CreatedBy     int        `json:"created_by"`
	UpdatedAt     time.Time  `json:"updated_at"`
	UpdatedBy     int        `json:"updated_by"`
}

func GenerateBusinessResponse(business *model.Business) *BusinessResponse {
	return &BusinessResponse{
		Id:            business.Id,
		OwnerUserId:   business.OwnerUserId,
		Email:         business.Email,
		PhoneNumber:   business.PhoneNumber,
		Name:          business.Name,
		VerifiedAt:    business.VerifiedAt,
		DeactivatedAt: business.DeactivatedAt,
		CreatedAt:     business.CreatedAt,
		CreatedBy:     business.CreatedBy,
		UpdatedAt:     business.UpdatedAt,
		UpdatedBy:     business.UpdatedBy,
	}
}
