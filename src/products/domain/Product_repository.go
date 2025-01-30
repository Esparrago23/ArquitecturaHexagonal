package domain

import (
	"demo/src/products/domain/entities"
)

type IProduct interface {
	Save(product *entities.Product)
	GetAll() ([]entities.Product)
	Delete(id int)
	Edit(id int,updatedProduct *entities.Product)
}