package controllers

import (
	"demo/src/products/application"
	"demo/src/products/domain/entities"
	"net/http"
    "time"
	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	useCase application.CreateProduct
}

func NewCreateProductController(useCase application.CreateProduct) *CreateProductController {
	return &CreateProductController{useCase: useCase}
}

func (cp_c *CreateProductController) Execute(c *gin.Context) {
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

	newProduct :=entities.Product{
		Name: product.Name,
		Price: product.Price,
        Quantity: product.Quantity,
        Created_at: product.Created_at,
	}

    err := cp_c.useCase.Execute(&newProduct)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
   
}