package dto

import (
	"ocean-pos/internal/model"
	"time"
)

type UserRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type UserResponse struct {
	Id                    int        `json:"id"`
	Email                 string     `json:"email"`
	Name                  string     `json:"name"`
	PhoneNumber           string     `json:"phone_number"`
	IsEmailVerified       bool       `json:"is_email_verified"`
	EmailVerifiedAt       *time.Time `json:"email_verified_at"`
	IsPhoneNumberVerified bool       `json:"is_phone_number_verified"`
	PhoneNumberVerifiedAt *time.Time `json:"phone_verified_at"`
	DeactivatedAt         *time.Time `json:"deactivated_at"`
	LastLogin             *time.Time `json:"last_login"`
	CreatedAt             time.Time  `json:"created_at"`
	CreatedBy             int        `json:"created_by"`
	UpdatedAt             time.Time  `json:"updated_at"`
	UpdatedBy             int        `json:"updated_by"`
}

type UserListResponse []*UserResponse

func GenerateUserResponse(user *model.User) *UserResponse {
	return &UserResponse{
		Id:                    user.Id,
		Email:                 user.Email,
		Name:                  user.Name,
		PhoneNumber:           user.PhoneNumber,
		IsEmailVerified:       user.IsEmailVerified,
		EmailVerifiedAt:       user.EmailVerifiedAt,
		IsPhoneNumberVerified: user.IsPhoneNumberVerified,
		PhoneNumberVerifiedAt: user.PhoneNumberVerifiedAt,
		DeactivatedAt:         user.DeactivatedAt,
		LastLogin:             user.LastLogin,
		CreatedAt:             user.CreatedAt,
		CreatedBy:             user.CreatedBy,
		UpdatedAt:             user.UpdatedAt,
		UpdatedBy:             user.UpdatedBy,
	}
}

func GenerateUserListResponse(user model.UserList) *UserListResponse {
	userList := UserListResponse{}
	for _, data := range user {
		userResponse := GenerateUserResponse(&data)
		userList = append(userList, userResponse)
	}
	return &userList
}
