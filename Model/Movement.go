package model

import "time"

type Movement struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"`
	Type      string    `gorm:"not null"`
	Quantity  uint64    `gorm:"not null"`
	ProductId uint64    `gorm:"not null"`
	Reason    string    `gorm:"default:null"`
	Product   *Product  `gorm:"foreignKey:ProductId;references:Id;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
