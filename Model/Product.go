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
	Movements   []Movement `gorm:"foreignKey:ProductId"`
	Unitmeasure string     `gorm:"not null"`
	Suppliers   *Suppliers `gorm:"foreignKey:SuppliersId;references:Id"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
}
