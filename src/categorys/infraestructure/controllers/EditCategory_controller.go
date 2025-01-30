package controllers

import (
    "demo/src/categorys/application"
    "demo/src/categorys/domain/entities"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type EditCategoryController struct {
    useCase application.EditCategory
}

func NewEditCategoryController(useCase application.EditCategory) *EditCategoryController {
    return &EditCategoryController{useCase: useCase}
}

func (cc_c *EditCategoryController) Execute(c *gin.Context) {
    var category struct {
        Name        string `json:"name"`
        Description string `json:"description"`
    }
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    categoryID := c.Param("id")
    updatedCategory := entities.Category{
        Name:        category.Name,
        Description: category.Description,
    }

    id, err := strconv.Atoi(categoryID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
        return
    }
    err = cc_c.useCase.Execute(id, &updatedCategory)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}