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
	CheckMissingProductsService := application.NewMissingProduct(ps)
	CheckNewProductsService := application.NewCheckNewProducts(ps)

	createProductController := controllers.NewCreateProductController(*createProductService)
	viewProductController := controllers.NewListProductController(*viewProductService)
	EditProductController := controllers.NewEditProductController(*EditProductService)
	deleteProductController := controllers.NewDeleteProductController(*deleteProductService)
	CheckMissingProductController := controllers.NewMissingProductsController(*CheckMissingProductsService)
	CheckNewProductsController := controllers.NewCheckNewProductsController(*CheckNewProductsService)

	ProductsRoutes(router, ProductHandlers{
		Create: createProductController,
		Get: viewProductController,
		Edit: EditProductController,
		Delete: deleteProductController,
		Missing: CheckMissingProductController,
		NewProducts: CheckNewProductsController,
	})
}