package application

import (
	"demo/src/products/domain"
	"demo/src/products/domain/entities"

)

type ViewProduct struct {
	db domain.IProduct
}

func NewViewProduct(db domain.IProduct) *ViewProduct {
	return &ViewProduct{db: db}
}

func (vp *ViewProduct) Execute() ([]entities.Product) {
    products := vp.db.GetAll()
    return products
}