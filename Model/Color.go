package model

import (
	"time"
)

type Color struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"not null"`
	TotalPieces uint64    `gorm:"not null"`
	TotalPrice  float64   `gorm:"not null"`
	Average     string    `gorm:"default:null"`
	CutOrderId  uint64    `gorm:"not null"`
	CutSizes    []CutSize `gorm:"foreignKey:ColorId;references:Id;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	CutOrder    *CutOrder `gorm:"foreignKey:CutOrderId;references:Id;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
