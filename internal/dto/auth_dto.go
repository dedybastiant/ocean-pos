package dto

import (
	"ocean-pos/internal/model"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateLoginResponse(auth model.Auth) LoginResponse {
	return LoginResponse{
		Token:        auth.Token,
		RefreshToken: auth.RefreshToken,
	}
}
