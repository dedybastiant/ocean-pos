package model

import "time"

type User struct {
	Id                    int
	Email                 string
	Password              string
	Name                  string
	PhoneNumber           string
	IsEmailVerified       bool
	EmailVerifiedAt       *time.Time
	IsPhoneNumberVerified bool
	PhoneNumberVerifiedAt *time.Time
	DeactivatedAt         *time.Time
	LastLogin             *time.Time
	CreatedAt             time.Time
	CreatedBy             int
	UpdatedAt             time.Time
	UpdatedBy             int
}

type UserList []User
