package model

import (
	"time"
)

type CutMovements struct {
	Id         uint64    `gorm:"primaryKey;autoIncrement"`
	MovementId uint64    `gorm:"not null"`
	CutOrderId uint64    `gorm:"not null"`
	Movement   *Movement `gorm:"foreignKey:MovementId;references:Id;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	CutOrder   *CutOrder `gorm:"foreignKey:CutOrderId;references:Id;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
