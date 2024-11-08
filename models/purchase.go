package models

import (
	"gorm.io/gorm"
	"time"
)

type Purchase struct {
	gorm.Model
	Date       time.Time         `json:"date"`
	Supplier   *Supplier         `json:"supplier"`
	SupplierID uint64            `json:"supplier_id"`
	TotalUnits int               `json:"total_units"`
	TotalPrice uint64            `json:"total_price"`
	Details    []*PurchaseDetail `json:"details"`
}
