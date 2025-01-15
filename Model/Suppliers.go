package model

import (
	"time"
)

type Suppliers struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"not null"`
	Contact   string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
	Phone     string    `gorm:"not null"`
	Address   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
