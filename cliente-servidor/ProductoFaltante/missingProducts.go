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
	Id int32
	Name string
	Price float32
	Quantity int32
}
func CheckingMissingProducts(c *gin.Context) {

	for {
		resp, err := http.Get(`http://localhost:8080/products/missing`)

		if err != nil {
            fmt.Printf("Error con petición: %s\n", err.Error())
            
        }

        if resp.StatusCode != http.StatusOK {
            fmt.Printf("Respuesta inesperada: %d\n", resp.StatusCode)
            
        }
		contentType := resp.Header.Get("Content-Type")
		if !strings.HasPrefix (contentType,"application/json") {
			fmt.Println("Respuesta inesperada del servidor, tipo de contenido:", contentType)
			
		}
		var response struct {
			Products []Product `json:"products"`
		}

		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			fmt.Printf("Error al decodificar la respuesta: %s\n", err.Error())
			
		}

		products := response.Products

        fmt.Println("Productos faltantes:")
        for _, product := range products {
            fmt.Printf("ID: %d, Nombre: %s, Precio: %.2f, Stock: %d\n", product.Id, product.Name, product.Price, product.Quantity)
        }


		time.Sleep(10 * time.Second)
	}

}