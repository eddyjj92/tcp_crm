package controllers

import "tcp_crm/models"

type ProductController struct{}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (c *ProductController) List() []models.Product {
	var products []models.Product
	DB.Find(&products)
	return products
}
