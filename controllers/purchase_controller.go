package controllers

import (
	"fmt"
	"tcp_crm/models"
)

type PurchaseController struct{}

func NewPurchaseController() *PurchaseController {
	return &PurchaseController{}
}

func (c *PurchaseController) List() ([]models.Purchase, error) {
	var purchases []models.Purchase
	// Preload de relaciones necesarias, puedes agregar más según tus modelos
	result := DB.Preload("Supplier").Preload("Details").Find(&purchases)
	// Manejo de errores
	if result.Error != nil {
		return nil, result.Error
	}
	return purchases, nil
}

func (c *PurchaseController) Store(purchase models.Purchase) (*models.Purchase, error) {
	// Validar el producto antes de guardarlo
	if err := purchase.Validate(); err != nil {
		return nil, err
	}
	DB.Create(&purchase)
	return &purchase, nil
}

func (c *PurchaseController) Update(id int, purchase models.Purchase) (*models.Purchase, error) {
	if err := purchase.Validate(); err != nil {
		return nil, err
	}

	edit := &models.Purchase{}
	if err := DB.First(edit, id).Error; err != nil {
		return nil, fmt.Errorf("compra no encontrads: %w", err)
	}

	// Actualizar los campos
	edit.Date = purchase.Date
	edit.SupplierID = purchase.SupplierID
	edit.TotalUnits = purchase.TotalUnits
	edit.TotalPrice = purchase.TotalPrice

	// Guardar los cambios en la base de datos
	if err := DB.Save(edit).Error; err != nil {
		return nil, fmt.Errorf("error al actualizar la compra: %w", err)
	}

	return edit, nil
}

func (c *PurchaseController) Delete(id int) (*models.Purchase, error) {
	// Buscar el producto existente para obtener los datos actuales
	remove := &models.Purchase{}
	if err := DB.First(remove, id).Error; err != nil {
		return nil, fmt.Errorf("compra no encontrada: %w", err)
	}
	DB.Delete(remove)
	return remove, nil
}
