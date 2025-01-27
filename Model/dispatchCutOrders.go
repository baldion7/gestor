package model

import "time"

type DispatchCutOrders struct {
	ID        int       `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
