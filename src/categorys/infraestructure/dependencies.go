package infraestructure

import (
    "demo/src/categorys/application"
    "demo/src/categorys/infraestructure/controllers"

    "github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {

    cs := NewMySQL()

    createCategoryService := application.NewCreateCategory(cs)
    viewCategoryService := application.NewViewCategory(cs)
    editCategoryService := application.NewEditCategory(cs)
    deleteCategoryService := application.NewDeleteCategory(cs)

    createCategoryController := controllers.NewCreateCategoryController(*createCategoryService)
    viewCategoryController := controllers.NewListCategoryController(*viewCategoryService)
    editCategoryController := controllers.NewEditCategoryController(*editCategoryService)
    deleteCategoryController := controllers.NewDeleteCategoryController(*deleteCategoryService)

    CategoryRoutes(router, CategoryHandlers{
        Create: createCategoryController,
        Get:    viewCategoryController,
        Edit:   editCategoryController,
        Delete: deleteCategoryController,
    })
}