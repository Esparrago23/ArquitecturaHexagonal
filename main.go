package main

import (
	"demo/src/products/infraestructure"
    categoryInfra "demo/src/categorys/infraestructure"
	"github.com/gin-gonic/gin"
)
	

func main() {
	r := gin.Default()
	infraestructure.Init(r)
	categoryInfra.Init(r)
	r.Run(":8080")
	
}