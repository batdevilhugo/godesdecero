package main

import (
	"fmt"

	"github.com/godesdecero/funciones"
	"github.com/godesdecero/variables"
)

func main() {
	fmt.Println("Hello Go")
	fmt.Println("Golng is my favorite language")

	// Llamando a mi primera funcion
	variables.Enteras()
	
	estado, texto := funciones.ConvierteTexto(15)
	fmt.Println(estado, texto)
}