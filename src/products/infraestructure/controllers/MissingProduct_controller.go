package controllers

import (
	"demo/src/products/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MissingProductsController struct {
	useCase application.MissingProducts
}

func NewMissingProductsController(useCase application.MissingProducts) *MissingProductsController {
	return &MissingProductsController{useCase: useCase}
}

func (mp_c *MissingProductsController) Execute(c *gin.Context) {
    products, err := mp_c.useCase.Execute()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"products": products})
}