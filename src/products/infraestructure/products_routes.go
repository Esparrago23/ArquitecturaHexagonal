package infraestructure

import (
	"demo/src/products/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

type ProductHandlers struct {
	Create *controllers.CreateProductController
	Get    *controllers.ListProductController
	Edit *controllers.EditProductController
	Delete *controllers.DeleteProductController
	Missing *controllers.MissingProductsController
	NewProducts *controllers.CheckNewProductsController
}
func ProductsRoutes(router *gin.Engine,handlers ProductHandlers){
	productsGroup := router.Group("/products")
    {
        productsGroup.POST("/", handlers.Create.Execute)
        productsGroup.GET("/", handlers.Get.Execute)
        productsGroup.PUT("/:id", handlers.Edit.Execute)
        productsGroup.DELETE("/:id", handlers.Delete.Execute)
		productsGroup.GET("/missing", handlers.Missing.Execute)
		productsGroup.GET("/new", handlers.NewProducts.Execute)
    }
}
