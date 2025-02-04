package main

import (
	"time";
	
	"demo/ProductoFaltante"
)

func main() {

	go ProductoFaltante.Run()
	time.Sleep(10 * time.Minute)
}
