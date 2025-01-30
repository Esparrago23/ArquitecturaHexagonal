package controllers

import (
    "demo/src/categorys/application"
    "demo/src/categorys/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

type CreateCategoryController struct {
    useCase application.CreateCategory
}

func NewCreateCategoryController(useCase application.CreateCategory) *CreateCategoryController {
    return &CreateCategoryController{useCase: useCase}
}

func (cc_c *CreateCategoryController) Execute(c *gin.Context) {
    var category struct {
        Name        string `json:"name"`
        Description string `json:"description"`
    }
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newCategory := entities.Category{
        Name:        category.Name,
        Description: category.Description,
    }

    err := cc_c.useCase.Execute(&newCategory)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Category created successfully"})
}