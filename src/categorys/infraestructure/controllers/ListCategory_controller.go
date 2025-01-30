package controllers

import (
    "demo/src/categorys/application"
    "net/http"

    "github.com/gin-gonic/gin"
)

type ListCategoryController struct {
    useCase application.ViewCategory
}

func NewListCategoryController(useCase application.ViewCategory) *ListCategoryController {
    return &ListCategoryController{useCase: useCase}
}

func (vc_c *ListCategoryController) Execute(c *gin.Context) {
    categories, err := vc_c.useCase.Execute()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"categories": categories})
}