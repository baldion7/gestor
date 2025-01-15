package model

import "time"

type CutOrder struct {
	Id            uint64  `gorm:"primaryKey;autoIncrement"`
	CreateBy      string  `gorm:"not null"`
	Quality       bool    `gorm:"not null"`
	Arrival       bool    `gorm:"not null"`
	Delivered     bool    `gorm:"not null"`
	TotalPieces   uint64  `gorm:"not null"`
	PricePerPiece float64 `gorm:"not null"`
	TotalPrice    float64 `gorm:"not null"`
	Observations  string  `gorm:"not null"`
	ReferenceId   uint64  `gorm:"not null"`
	CarvingsId    uint64
	Carvings      *Carvings  `gorm:"foreignKey:CarvingsId;references:Id"`
	Reference     *Reference `gorm:"foreignKey:ReferenceId;references:Id"`
	CreatedAt     time.Time  `gorm:"autoCreateTime"`
	UpdatedAt     time.Time  `gorm:"autoUpdateTime"`
}
