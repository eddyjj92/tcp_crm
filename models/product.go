package models

import (
	"tcp_crm/helpers"
	"time"
)

// Estructura del Producto
type Product struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `json:"name" validate:"required,min=3,max=100" alias:"Nombre"`
	Description   string    `json:"description" validate:"required,max=255" alias:"Descripción"`
	PurchasePrice string    `json:"purchase_price" validate:"required,gt=0,numeric" alias:"Precio de compra"`
	SalePrice     string    `json:"sale_price" validate:"required,gt=0,numeric" alias:"Precio de venta"`
	ImagePath     string    `json:"image_path" validate:"omitempty" alias:"Ruta de la imagen"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Implementar el método Validate para el modelo Product
func (p *Product) Validate() error {
	aliases := map[string]string{
		"Name":             "Nombre",
		"Description":      "Descripción",
		"PurchasePrice":    "Precio de compra",
		"SalePrice":        "Precio de venta",
		"ProfitPercentage": "Porcentaje de ganancia",
		"ImagePath":        "Ruta de la imagen",
	}
	return helpers.ValidateStruct(p, aliases)
}
