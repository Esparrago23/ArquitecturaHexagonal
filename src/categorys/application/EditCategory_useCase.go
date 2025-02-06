package application

import (
    "demo/src/categorys/domain"
    "demo/src/categorys/domain/entities"
)

type EditCategory struct {
    db domain.ICategory
}

func NewEditCategory(db domain.ICategory) *EditCategory {
    return &EditCategory{db: db}
}

func (ec *EditCategory) Execute(id int, updatedCategory *entities.Category) error {
    err := ec.db.Edit(id, updatedCategory)
    if err != nil {
        return err
    }
    return nil
}