package model

import "time"

type Business struct {
	Id            int
	OwnerUserId   int
	Email         string
	PhoneNumber   string
	Name          string
	VerifiedAt    *time.Time
	DeactivatedAt *time.Time
	CreatedAt     time.Time
	CreatedBy     int
	UpdatedAt     time.Time
	UpdatedBy     int
}
