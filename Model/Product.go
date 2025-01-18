package model

import (
	"time"
)

type Product struct {
	Id          uint64     `gorm:"primaryKey;autoIncrement"`
	Name        string     `gorm:"not null"`
	Reference   string     `gorm:"not null"`
	Color       string     `gorm:"not null"`
	Size        string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	Price       float64    `gorm:"not null"`
	SuppliersId uint64     `gorm:"not null"`
	Unitmeasure string     `gorm:"not null"`
	Movements   []Movement `gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	Suppliers   *Suppliers `gorm:"foreignKey:SuppliersId;references:Id;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
}
