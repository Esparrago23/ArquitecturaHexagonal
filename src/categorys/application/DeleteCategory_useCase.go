package application

import (
    "demo/src/categorys/domain"
)

type DeleteCategory struct {
    db domain.ICategory
}

func NewDeleteCategory(db domain.ICategory) *DeleteCategory {
    return &DeleteCategory{db: db}
}

func (dc *DeleteCategory) Execute(id int) error {
    err := dc.db.Delete(id)
    if err != nil {
        return err
    }
    return nil
}