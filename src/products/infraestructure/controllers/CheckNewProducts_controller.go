package controllers

import (
	"demo/src/products/domain/entities"
	"demo/src/products/application"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CheckNewProductsController struct {
	useCase application.CheckNewProducts
}

func NewCheckNewProductsController(useCase application.CheckNewProducts) *CheckNewProductsController {
	return &CheckNewProductsController{useCase: useCase}
}

func (mp_c *CheckNewProductsController) Execute(c *gin.Context) {
    productsChan := make(chan []entities.Product)
	errorChan := make(chan error)
	timeout := time.NewTimer(30 * time.Second)
	defer timeout.Stop()


	go func() {
		ticker := time.NewTicker(20 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			products, err := mp_c.useCase.Execute()
			if err != nil {
				errorChan <- err
				return
			}
			if len(products) > 0 {
				productsChan <- products
				return
			}
		}
	}()

	select {
	case products := <-productsChan:
		c.JSON(http.StatusOK, gin.H{"products": products})

	case err := <-errorChan:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	case <-timeout.C:
		c.JSON(http.StatusNoContent, gin.H{"message": "No hay nuevos productos"})
	}
}