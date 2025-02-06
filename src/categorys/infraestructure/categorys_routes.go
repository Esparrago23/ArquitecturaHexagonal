package infraestructure

import (
    "demo/src/categorys/infraestructure/controllers"

    "github.com/gin-gonic/gin"
)

type CategoryHandlers struct {
    Create *controllers.CreateCategoryController
    Get    *controllers.ListCategoryController
    Edit   *controllers.EditCategoryController
    Delete *controllers.DeleteCategoryController
}

func CategoryRoutes(router *gin.Engine, handlers CategoryHandlers) {
    categoriesGroup := router.Group("/categories")
    {
        categoriesGroup.POST("/", handlers.Create.Execute)
        categoriesGroup.GET("/", handlers.Get.Execute)
        categoriesGroup.PUT("/:id", handlers.Edit.Execute)
        categoriesGroup.DELETE("/:id", handlers.Delete.Execute)
    }
}