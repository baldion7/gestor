package model

import "time"

type CutSize struct {
	Id              uint64    `gorm:"primaryKey;autoIncrement"`
	Size            string    `gorm:"not null"`
	Quantity        uint64    `gorm:"not null"`
	ArrivalQuantity uint64    `gorm:"not null"`
	ColorId         uint64    `gorm:"not null"`
	Color           *Color    `gorm:"foreignKey:ColorId;references:Id"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
