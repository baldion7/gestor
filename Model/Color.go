package model

import (
	"time"
)

type Color struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"not null"`
	TotalPieces uint64    `gorm:"not null"`
	TotalPrice  float64   `gorm:"not null"`
	CutOrderId  uint64    `gorm:"not null"`
	CutOrder    *CutOrder `gorm:"foreignKey:CutOrderId;references:Id"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
