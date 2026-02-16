package main

import (
	"fmt"

	"github.com/godesdecero/condicionales"
	ejercicio "github.com/godesdecero/ejercicios"
	"github.com/godesdecero/funciones"
	"github.com/godesdecero/teclado"
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

	// Ejercicio1
	estado3, numero := ejercicio.Ejercicio01("3") // archivo.funcion
	fmt.Println(estado3, numero)

	// Llamando a teclado
	teclado.IngresaNumero()
}