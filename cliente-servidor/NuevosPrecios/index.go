package NuevosPrecios

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("/price-changes", CheckPriceChanges)

	server := &http.Server{
		Addr:         ":4003",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 5 * time.Minute,
		IdleTimeout:  1 * time.Hour,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error: No se pudo iniciar el servidor")
	}
}
