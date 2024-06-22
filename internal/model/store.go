package model

import "time"

type Store struct {
	Id            int
	BusinessId    int
	Name          string
	Location      string
	Description   string
	DeactivatedAt *time.Time
	CreatedAt     time.Time
	CreatedBy     int
	UpdatedAt     time.Time
	UpdatedBy     int
}
