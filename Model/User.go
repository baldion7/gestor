package model

import (
	"time"
)

type User struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	RoleId    uint64    `gorm:"not null"`
	Role      *Role     `gorm:"foreignKey:RoleId;references:Id;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
