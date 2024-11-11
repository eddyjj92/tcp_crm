package controllers

import (
	"fmt"
	"tcp_crm/helpers"
	"tcp_crm/models"
)

type ProductController struct{}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (c *ProductController) List() []models.Product {
	var products []models.Product
	DB.Find(&products)
	return products
}

func (c *ProductController) Store(product models.Product) (*models.Product, error) {
	// Validar el producto antes de guardarlo
	if err := product.Validate(); err != nil {
		return nil, err
	}
	if product.ImagePath != "" {
		fileName, err := helpers.UploadImage(product.ImagePath)
		if err != nil {
			return nil, err
		}
		product.ImagePath = fmt.Sprintf("products/%s", fileName)
	} else {
		product.ImagePath = fmt.Sprintf("products/product.png")
	}
	DB.Create(&product)
	return &product, nil
}

func (c *ProductController) Update(id int, product models.Product) (*models.Product, error) {
	// Validar el producto antes de actualizarlo
	if err := product.Validate(); err != nil {
		return nil, err
	}
	// Buscar el producto existente para obtener los datos actuales
	existingProduct := &models.Product{}
	if err := DB.First(existingProduct, id).Error; err != nil {
		return nil, fmt.Errorf("producto no encontrado: %w", err)
	}
	// Subir imagen si existe una nueva imagen proporcionada
	if product.ImagePath != "" {
		fileName, err := helpers.UploadImage(product.ImagePath)
		if err != nil {
			return nil, err
		}
		existingProduct.ImagePath = fmt.Sprintf("products/%s", fileName)
	}
	// Actualizar los campos del producto existente
	existingProduct.Name = product.Name
	existingProduct.Description = product.Description
	existingProduct.PurchasePrice = product.PurchasePrice
	existingProduct.SalePrice = product.SalePrice

	// Guardar los cambios en la base de datos
	if err := DB.Save(existingProduct).Error; err != nil {
		return nil, fmt.Errorf("error al actualizar el producto: %w", err)
	}

	return existingProduct, nil
}

func (c *ProductController) Delete(id int) (*models.Product, error) {
	// Buscar el producto existente para obtener los datos actuales
	remove := &models.Product{}
	if err := DB.First(remove, id).Error; err != nil {
		return nil, fmt.Errorf("producto no encontrado: %w", err)
	}
	DB.Delete(remove)
	return remove, nil
}
