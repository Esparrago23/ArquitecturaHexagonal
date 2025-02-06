package NuevoProducto

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	s := gin.Default()
	s.GET("/new/", CheckingNewProducts)

	srv := &http.Server{
		Addr:         ":4002",
		Handler:      s,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 5 * time.Minute,
		IdleTimeout:  1 * time.Hour,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Error: Server hasn't started")
	}
}