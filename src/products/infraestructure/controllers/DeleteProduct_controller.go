package controllers

import (
	"demo/src/products/application"
	
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	useCase application.DeleteProduct
}

func NewDeleteProductController(useCase application.DeleteProduct) *DeleteProductController {
	return &DeleteProductController{useCase: useCase}
}

func (dp_c *DeleteProductController) Execute(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.Atoi(productID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    err = dp_c.useCase.Execute(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product Deleted successfully"})
   
}