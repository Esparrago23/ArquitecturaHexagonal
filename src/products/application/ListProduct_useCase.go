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

func (vp *ViewProduct) Execute() ([]entities.Product,error) {
    res, err := vp.db.GetAll()
	if err != nil {
		return res,err
	}
	return res,nil
}