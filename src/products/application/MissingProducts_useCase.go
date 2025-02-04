package application

import (
	"demo/src/products/domain"
	"demo/src/products/domain/entities"
)

type MissingProducts struct {
	db domain.IProduct
}

func NewMissingProduct(db domain.IProduct) *MissingProducts {
	return &MissingProducts{db: db}
}
func (mp *MissingProducts) Execute()([]entities.Product,error){
	res, err := mp.db.CheckMissingProducts()
	if err != nil {
		return res,err
	}
	return res,nil
}
