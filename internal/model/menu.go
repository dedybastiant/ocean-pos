package model

import "time"

type Menu struct {
	Id            int
	CategoryId    int
	Name          string
	DefaultPrice  int
	DeactivatedAt *time.Time
	CreatedAt     time.Time
	CreatedBy     int
	UpdatedAt     time.Time
	UpdatedBy     int
}
