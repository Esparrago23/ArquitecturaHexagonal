package ProductoFaltante

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Do_after_init(c *gin.Context) {
	CheckingMissingProducts(c)
}


func Run() {
	s := gin.Default()

	s.GET("/missing/", CheckingMissingProducts)

	srv2 := &http.Server{
		Addr:         ":4001",
		Handler:      s,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 5 * time.Minute,
		IdleTimeout:  1 * time.Hour,
	}

	if err := srv2.ListenAndServe(); err != nil {
		fmt.Println("Error: Server hasn't begin")
	}
}
