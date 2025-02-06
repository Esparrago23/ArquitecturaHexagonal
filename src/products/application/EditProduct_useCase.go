package application

import (
	"demo/src/products/domain"
	"demo/src/products/domain/entities"
)

type EditProduct struct {
	db domain.IProduct
}

func NewEditProduct(db domain.IProduct) *EditProduct {
	return &EditProduct{db: db}
}

func (ep *EditProduct) Execute(id int,updatedProduct *entities.Product ) error {
	err:=ep.db.Edit(id,updatedProduct)
	if err != nil {
		return err
	}
	return nil
}