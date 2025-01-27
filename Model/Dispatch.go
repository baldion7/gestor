package model

import (
	"gorm.io/gorm"
	"time"
)

type Dispatch struct {
	Id             uint64      `gorm:"primaryKey;autoIncrement"`
	BrandId        uint64      `gorm:"not null;index"`
	DispatchNumber uint64      `gorm:"not null"`
	TotalBag       uint64      `gorm:"not null"`
	Collect        string      `gorm:"not null"`
	Delivery       string      `gorm:"not null"`
	Boxes          uint64      `gorm:"not null"`
	Brand          *Brand      `gorm:"foreignKey:BrandId;references:Id;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	CreatedAt      time.Time   `gorm:"autoCreateTime"`
	UpdatedAt      time.Time   `gorm:"autoUpdateTime"`
	CutOrders      []*CutOrder `gorm:"many2many:dispatch_cut_orders;"`
}

func (d *Dispatch) BeforeCreate(tx *gorm.DB) (err error) {
	// Find the last dispatch number for this brand
	var lastDispatch Dispatch
	result := tx.Where("brand_id = ?", d.BrandId).
		Order("dispatch_number DESC").
		First(&lastDispatch)

	if result.Error == gorm.ErrRecordNotFound {
		// First dispatch for this brand
		d.DispatchNumber = 1
	} else {
		// Increment last dispatch number
		d.DispatchNumber = lastDispatch.DispatchNumber + 1
	}

	return nil
}
