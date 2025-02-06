package application

import (
	"demo/src/products/domain"
	"demo/src/products/domain/entities"
)

type CheckNewProducts struct {
	db domain.IProduct
}

func NewCheckNewProducts(db domain.IProduct) *CheckNewProducts {
	return &CheckNewProducts{db: db}
}
func (mp *CheckNewProducts) Execute()([]entities.Product,error){
	res, err := mp.db.CheckNewProducts()
	if err != nil {
		return res,err
	}
	return res,nil
}
