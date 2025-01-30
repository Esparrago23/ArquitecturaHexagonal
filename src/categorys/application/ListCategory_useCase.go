package application

import (
    "demo/src/categorys/domain"
    "demo/src/categorys/domain/entities"
)

type ViewCategory struct {
    db domain.ICategory
}

func NewViewCategory(db domain.ICategory) *ViewCategory {
    return &ViewCategory{db: db}
}

func (vc *ViewCategory) Execute() ([]entities.Category, error) {
    res, err := vc.db.GetAll()
    if err != nil {
        return res, err
    }
    return res, nil
}