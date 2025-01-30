package controllers

import (
	"demo/src/products/application"
	"demo/src/products/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EditProductController struct {
	useCase application.EditProduct
}

func NewEditProductController(useCase application.EditProduct) *EditProductController {
	return &EditProductController{useCase: useCase}
}

func (cp_c *EditProductController) Execute(c *gin.Context) {
    var product struct {
        Name  string  `json:"name"`
        Price float32 `json:"price"`
    }
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    productID := c.Param("id")
    updatedProduct := entities.Product{
        Name:  product.Name,
        Price: product.Price,
    }

    id, err := strconv.Atoi(productID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }
    err = cp_c.useCase.Execute(id, &updatedProduct)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}