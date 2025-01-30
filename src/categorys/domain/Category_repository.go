package domain

import (
    "demo/src/categorys/domain/entities"
)

type ICategory interface {
    Save(category *entities.Category) error
    GetAll() ([]entities.Category, error)
    Delete(id int) error
    Edit(id int, updatedCategory *entities.Category) error
}