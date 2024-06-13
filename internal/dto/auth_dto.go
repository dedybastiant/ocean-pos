package dto

import (
	"ocean-pos/internal/model"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateAuthResponse(auth model.Auth) AuthResponse {
	return AuthResponse{
		Token:        auth.Token,
		RefreshToken: auth.RefreshToken,
	}
}
