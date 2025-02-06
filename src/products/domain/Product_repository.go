package domain

import (
	"demo/src/products/domain/entities"
)

type IProduct interface {
	Save(product *entities.Product) error
	GetAll() ([]entities.Product, error)
	Delete(id int) error
	Edit(id int,updatedProduct *entities.Product) error
	CheckMissingProducts()([]entities.Product, error)
	CheckNewProducts()([]entities.Product, error)

}