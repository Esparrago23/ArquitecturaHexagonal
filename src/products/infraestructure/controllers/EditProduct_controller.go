package controllers

import (
	"demo/src/products/application"
	"demo/src/products/domain/entities"
	"net/http"
	"strconv"
    "time"
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
        Quantity int32   `json:"quantity"`
        Created_at     time.Time `json:"created_at"`
    }
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if product.Price < 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be zero or greater"})
        return
    }
    if product.Quantity < 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity must be zero or greater"})
        return
    }


    productID := c.Param("id")
    updatedProduct := entities.Product{
        Name:  product.Name,
        Price: product.Price,
        Quantity: product.Quantity,
        Created_at: product.Created_at,
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