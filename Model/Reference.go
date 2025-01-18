package model

type Reference struct {
	Id                uint64  `gorm:"primaryKey;autoIncrement"`
	Name              string  `gorm:"not null"`
	BrandId           uint64  `gorm:"not null"`
	CostPerProduction float64 `gorm:"not null"`
	EnsemblePrice     float64 `gorm:"not null"`
	Brand             *Brand  `gorm:"foreignKey:BrandId;references:Id;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
}
