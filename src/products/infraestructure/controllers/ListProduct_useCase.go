package controllers

import (
	"demo/src/products/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListProductController struct {
	useCase application.ViewProduct
}

func NewListProductController(useCase application.ViewProduct) *ListProductController {
	return &ListProductController{useCase: useCase}
}

func (vp_c *ListProductController) Execute(c *gin.Context) {
    products, err := vp_c.useCase.Execute()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"products": products})
}