package ProductoFaltante

import (
	"fmt"
	"net/http"
	_"strconv"
	"time"
	"strings"
    "encoding/json"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	Quantity  int32     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}
var lastProducts map[int32]Product 

func CheckingMissingProducts(c *gin.Context) {
	for {
		resp, err := http.Get("http://localhost:8080/products/missing")
		if err != nil {
			fmt.Printf("Error con petición: %s\n", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Respuesta inesperada: %d\n", resp.StatusCode)
			time.Sleep(1 * time.Second)
			continue
		}

		contentType := resp.Header.Get("Content-Type")
		if !strings.HasPrefix(contentType, "application/json") {
			fmt.Println("Tipo de contenido inesperado:", contentType)
			time.Sleep(1 * time.Second)
			continue
		}

		var response struct {
			Products []Product `json:"products"`
		}

		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			fmt.Printf("Error al decodificar la respuesta: %s\n", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}

		products := filterLowStock(response.Products)

		newProducts := getNewProducts(products)
		removedProducts := getRemovedProducts(products)

		if len(newProducts) == 0 && len(removedProducts) == 0 {
			fmt.Println("No hay cambios en los productos con bajo stock.")
		} else {
			if len(newProducts) > 0 {
				fmt.Println("Nuevos productos con stock bajo:")
				for _, product := range newProducts {
					fmt.Printf("ID: %d, Nombre: %s, Precio: %.2f, Stock: %d\n", product.Id, product.Name, product.Price, product.Quantity)
				}
			}

			if len(removedProducts) > 0 {
				fmt.Println("Productos que ya no están en la lista de stock bajo:")
				for _, product := range removedProducts {
					fmt.Printf("ID: %d, Nombre: %s\n", product.Id, product.Name)
				}
			}
		}
		updateLastProducts(products)

		time.Sleep(5 * time.Second)
	}
}

func filterLowStock(products []Product) []Product {
	var filtered []Product
	for _, product := range products {
		if product.Quantity < 5 {
			filtered = append(filtered, product)
		}
	}
	return filtered
}

func getNewProducts(products []Product) []Product {
	var newProducts []Product
	for _, product := range products {
		if _, exists := lastProducts[product.Id]; !exists {
			newProducts = append(newProducts, product)
		}
	}
	return newProducts
}

func getRemovedProducts(products []Product) []Product {
	newProductMap := make(map[int32]bool)
	for _, product := range products {
		newProductMap[product.Id] = true
	}

	var removedProducts []Product
	for id, product := range lastProducts {
		if _, exists := newProductMap[id]; !exists {
			removedProducts = append(removedProducts, product)
		}
	}
	return removedProducts
}

func updateLastProducts(products []Product) {
	lastProducts = make(map[int32]Product)
	for _, product := range products {
		lastProducts[product.Id] = product
	}
}