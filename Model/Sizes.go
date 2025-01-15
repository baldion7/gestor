package model

import (
	"time"
)

type Sizes struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
