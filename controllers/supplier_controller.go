package controllers

import (
	"fmt"
	"tcp_crm/models"
)

type SupplierController struct{}

func NewSupplierController() *SupplierController {
	return &SupplierController{}
}

func (c *SupplierController) List() []models.Supplier {
	var suppliers []models.Supplier
	DB.Find(&suppliers)
	return suppliers
}

func (c *SupplierController) Store(supplier models.Supplier) (*models.Supplier, error) {
	// Validar el producto antes de guardarlo
	if err := supplier.Validate(); err != nil {
		return nil, err
	}
	DB.Create(&supplier)
	return &supplier, nil
}

func (c *SupplierController) Update(id int, supplier models.Supplier) (*models.Supplier, error) {
	if err := supplier.Validate(); err != nil {
		return nil, err
	}

	edit := &models.Supplier{}
	if err := DB.First(edit, id).Error; err != nil {
		return nil, fmt.Errorf("proveedor no encontrado: %w", err)
	}

	// Actualizar los campos
	edit.Name = supplier.Name
	edit.Address = supplier.Address
	edit.Phone = supplier.Phone
	edit.Email = supplier.Email

	// Guardar los cambios en la base de datos
	if err := DB.Save(edit).Error; err != nil {
		return nil, fmt.Errorf("error al actualizar el proveedor: %w", err)
	}

	return edit, nil
}

func (c *SupplierController) Delete(id int) (*models.Supplier, error) {
	// Buscar el producto existente para obtener los datos actuales
	remove := &models.Supplier{}
	if err := DB.First(remove, id).Error; err != nil {
		return nil, fmt.Errorf("proveedor no encontrado: %w", err)
	}
	DB.Delete(remove)
	return remove, nil
}
