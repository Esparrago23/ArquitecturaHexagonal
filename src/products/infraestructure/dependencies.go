package infraestructure

import (
	"demo/src/products/application"
	"demo/src/products/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)


func Init(router *gin.Engine) {
	
	
	ps := NewMySQL()

	createProductService := application.NewCreateProduct(ps)
	viewProductService := application.NewViewProduct(ps)
	EditProductService := application.NewEditProduct(ps)
	deleteProductService := application.NewDeleteProduct(ps)

	createProductController := controllers.NewCreateProductController(*createProductService)
	viewProductController := controllers.NewListProductController(*viewProductService)
	EditProductController := controllers.NewEditProductController(*EditProductService)
	deleteProductController := controllers.NewDeleteProductController(*deleteProductService)

	ProductsRoutes(router, ProductHandlers{
		Create: createProductController,
		Get: viewProductController,
		Edit: EditProductController,
		Delete: deleteProductController,
	})
}