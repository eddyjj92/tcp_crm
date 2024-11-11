package models

import (
	"time"
)

type PurchaseDetail struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Product       *Product  `json:"product"`
	ProductID     uint      `json:"product_id"`
	PurchasePrice float64   `json:"purchase_price"`
	SalePrice     float64   `json:"sale_price"`
	Units         int       `json:"units"`
	Subtotal      float64   `json:"subtotal"`
	Purchase      *Purchase `json:"purchase"`
	PurchaseID    uint      `json:"purchase_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
