package controllers

import (
    "demo/src/categorys/application"
    "strconv"
    "net/http"
    "github.com/gin-gonic/gin"
)

type DeleteCategoryController struct {
    useCase application.DeleteCategory
}

func NewDeleteCategoryController(useCase application.DeleteCategory) *DeleteCategoryController {
    return &DeleteCategoryController{useCase: useCase}
}

func (dc_c *DeleteCategoryController) Execute(c *gin.Context) {
    categoryID := c.Param("id")
    id, err := strconv.Atoi(categoryID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
        return
    }

    err = dc_c.useCase.Execute(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}