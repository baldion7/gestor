package model

import (
	"time"
)

type Carvings struct {
	Id                 uint64    `gorm:"primaryKey;autoIncrement"`
	Name               string    `gorm:"not null"`
	Contact            string    `gorm:"not null"`
	Email              string    `gorm:"not null"`
	Phone              string    `gorm:"not null"`
	Address            string    `gorm:"not null"`
	ProductionCapacity uint64    `gorm:"not null"`
	Delivery           float64   `gorm:"not null"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
}
