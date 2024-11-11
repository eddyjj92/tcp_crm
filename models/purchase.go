package models

import (
	"tcp_crm/helpers"
	"time"
)

type Purchase struct {
	ID         uint              `gorm:"primaryKey" json:"id"`
	Date       time.Time         `json:"date"`
	Supplier   *Supplier         `json:"supplier"`
	SupplierID uint64            `json:"supplier_id"`
	TotalUnits int               `json:"total_units"`
	TotalPrice uint64            `json:"total_price"`
	Details    []*PurchaseDetail `json:"details"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

// Implementar el método Validate
func (p *Purchase) Validate() error {
	aliases := map[string]string{
		"Date":       "Fecha",
		"SupplierID": "Proveedor",
		"TotalUnits": "Total de Unidades",
		"TotalPrice": "Precio Total",
	}
	return helpers.ValidateStruct(p, aliases)
}
