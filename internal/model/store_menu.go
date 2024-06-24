package model

import "time"

type StoreMenu struct {
	Id            int
	StoreId       int
	MenuId        int
	StorePrice    int
	IsAvailable   bool
	DeactivatedAt *time.Time
	CreatedAt     time.Time
	CreatedBy     int
	UpdatedAt     time.Time
	UpdatedBy     int
}
