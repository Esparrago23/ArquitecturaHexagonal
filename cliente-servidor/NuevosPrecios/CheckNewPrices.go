package NuevosPrecios

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	Quantity  int32     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

var lastPrices map[int32]float32

func CheckPriceChanges(c *gin.Context) {
	for {
		resp, err := http.Get("http://localhost:8080/products/")
		if err != nil {
			fmt.Printf("Error en la petición: %s\n", err.Error())
			time.Sleep(5 * time.Second)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Respuesta inesperada: %d\n", resp.StatusCode)
			time.Sleep(5 * time.Second)
			continue
		}

		contentType := resp.Header.Get("Content-Type")
		if !strings.HasPrefix(contentType, "application/json") {
			fmt.Println("Tipo de contenido inesperado:", contentType)
			time.Sleep(5 * time.Second)
			continue
		}

		var response struct {
			Products []Product `json:"products"`
		}

		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			fmt.Printf("Error al decodificar la respuesta: %s\n", err.Error())
			time.Sleep(5 * time.Second)
			continue
		}

		priceChanges := getLowerPriceChanges(response.Products)

		if len(priceChanges) > 0 {
			fmt.Println("Productos con reducción de precio detectados:")
			for _, product := range priceChanges {
				fmt.Printf("ID: %d, Nombre: %s, Precio anterior: %.2f, Precio actual: %.2f\n",
					product.Id, product.Name, lastPrices[product.Id], product.Price)
			}
		} else {
			fmt.Println("No hay cambios en los precios.")
		}

		updateLastPrices(response.Products)
		time.Sleep(10 * time.Second)
	}
}

func getLowerPriceChanges(products []Product) []Product {
	var changedProducts []Product
	for _, product := range products {
		if lastPrice, exists := lastPrices[product.Id]; exists {
			if product.Price < lastPrice {
				changedProducts = append(changedProducts, product)
			}
		}
	}
	return changedProducts
}

func updateLastPrices(products []Product) {
	lastPrices = make(map[int32]float32)
	for _, product := range products {
		lastPrices[product.Id] = product.Price
	}
}
