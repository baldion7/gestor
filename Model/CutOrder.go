package model

import "time"

type CutOrder struct {
	Id            uint64         `gorm:"primaryKey;autoIncrement"`
	CreateBy      string         `gorm:"not null"`
	Quality       bool           `gorm:"not null"`
	Arrival       bool           `gorm:"not null"`
	Delivered     bool           `gorm:"not null"`
	Finish        bool           `gorm:"default:false"`
	TotalPieces   uint64         `gorm:"not null"`
	PricePerPiece float64        `gorm:"not null"`
	TotalPrice    float64        `gorm:"not null"`
	Observations  string         `gorm:"not null"`
	Observations2 string         `gorm:"default:null"`
	ReferenceId   uint64         `gorm:"not null"`
	CarvingsId    uint64         `gorm:"default:null"` // Cambiado a puntero
	Colors        []Color        `gorm:"foreignKey:CutOrderId;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	CutMovements  []CutMovements `gorm:"foreignKey:CutOrderId;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	Carvings      *Carvings      `gorm:"foreignKey:CarvingsId;references:Id;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	Reference     *Reference     `gorm:"foreignKey:ReferenceId;references:Id;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	Dispatches    []*Dispatch    `gorm:"many2many:dispatch_cut_orders;" json:"dispatches"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
}
