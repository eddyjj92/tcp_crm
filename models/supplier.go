package models

import (
	"tcp_crm/helpers"
	"time"
)

type Supplier struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name" validate:"required,min=3,max=100"`
	Address   string    `json:"address" validate:"required,max=255"`
	Phone     string    `json:"phone" validate:"required,min=8,numeric"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Implementar el método Validate para el modelo Product
func (s *Supplier) Validate() error {
	aliases := map[string]string{
		"Name":    "Nombre",
		"Address": "Dirección",
		"Phone":   "Teléfono",
		"Email":   "Correo",
	}
	return helpers.ValidateStruct(s, aliases)
}
