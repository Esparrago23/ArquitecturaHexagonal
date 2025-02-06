package main

import (
	"time";
	"demo/NuevoProducto";
	"demo/ProductoFaltante"
	"demo/NuevosPrecios"
)

func main() {
	go NuevosPrecios.Run()
	go ProductoFaltante.Run()
	go NuevoProducto.Run()
	time.Sleep(10 * time.Minute)
}
