package models

import "gorm.io/gorm"

type PurchaseDetail struct {
	gorm.Model
	Product       *Product `json:"product"`
	ProductID     string   `json:"product_id"`
	PurchasePrice float64  `json:"purchase_price"`
	SalePrice     float64  `json:"sale_price"`
	Units         int      `json:"units"`
	Subtotal      float64  `json:"subtotal"`
}
