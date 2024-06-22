package model

import "time"

type Category struct {
	Id            int
	BusinessId    int
	Name          string
	DeactivatedAt *time.Time
	CreatedAt     time.Time
	CreatedBy     int
	UpdatedAt     time.Time
	UpdatedBy     int
}
