package NuevoProducto

import (
	"fmt"
	"net/http"
	"time"
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

func CheckingNewProducts(c *gin.Context) {
	productsChan := make(chan []Product)
	errorChan := make(chan error)

	go func() {
		for {
			resp, err := http.Get("http://localhost:8080/products/new")
			if err != nil {
				errorChan <- fmt.Errorf("error en la petición: %s", err.Error())
				time.Sleep(5 * time.Second)
				continue
			}

			if resp.StatusCode == http.StatusNoContent {
				fmt.Println("No hay nuevos productos, esperando más cambios...")
				resp.Body.Close()
				continue
			}

			if resp.StatusCode != http.StatusOK {
				errorChan <- fmt.Errorf("respuesta inesperada: %d", resp.StatusCode)
				resp.Body.Close()
				time.Sleep(5 * time.Second)
				continue
			}

			var response struct {
				Products []Product `json:"products"`
			}
			err = json.NewDecoder(resp.Body).Decode(&response)
			resp.Body.Close()
			if err != nil {
				errorChan <- fmt.Errorf("error al decodificar la respuesta: %s", err.Error())
				time.Sleep(5 * time.Second)
				continue
			}
			productsChan <- response.Products
		}
	}()

	for {
		select {
		case products := <-productsChan:
			newProducts := getNewProducts(products)
			if len(newProducts) == 0 {
				fmt.Println("No hay nuevos productos agregados.")
			} else {
				fmt.Println("Nuevos productos detectados:")
				for _, product := range newProducts {
					fmt.Printf("ID: %d, Nombre: %s, Precio: %.2f, Creado en: %s\n",
						product.Id, product.Name, product.Price, product.CreatedAt)
				}
			}
			updateLastProducts(products)

		case err := <-errorChan:
			fmt.Printf("Error: %s\n", err.Error())
		}
	}
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


func updateLastProducts(products []Product) {
	if lastProducts == nil {
		lastProducts = make(map[int32]Product) 
	}
	for _, product := range products {
		lastProducts[product.Id] = product
	}
}