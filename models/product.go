package models

import "time"

type Product struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Product          string    `json:"product"`
	Description      string    `json:"description"`
	PurchasePrice    float64   `json:"purchase_price"`
	SalePrice        float64   `json:"sale_price"`
	ProfitPercentage float64   `json:"profit_percentage"`
	ImagePath        string    `json:"image_path"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
