package application

import (
	"demo/src/products/domain"
	"demo/src/products/domain/entities"
)

type CreateProduct struct {
	db domain.IProduct
}

func NewCreateProduct(db domain.IProduct) *CreateProduct {
	return &CreateProduct{db: db}
}

func (cp *CreateProduct) Execute(NewProduct *entities.Product){
	cp.db.Save(NewProduct)
}