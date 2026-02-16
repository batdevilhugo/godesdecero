package main

import (
	"fmt"

	"github.com/godesdecero/condicionales"
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

	estado2, numero := funciones.ConvierteNumero("1")
	fmt.Println(estado2, numero)

	// Llamada condicionales
	condicionales.SistemaOperatico()
}