package application

import (
    "demo/src/categorys/domain"
    "demo/src/categorys/domain/entities"
)

type CreateCategory struct {
    db domain.ICategory
}

func NewCreateCategory(db domain.ICategory) *CreateCategory {
    return &CreateCategory{db: db}
}

func (cc *CreateCategory) Execute(NewCategory *entities.Category) error {
    err := cc.db.Save(NewCategory)
    if err != nil {
        return err
    }
    return nil
}