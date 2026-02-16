package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func IngresaNumero() {
	fmt.Println("Ingresa numero 1: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		n1, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(fmt.Sprintf("[!] Dato ingresado incorrectamente", err.Error()))
		} else {
			fmt.Println("Numero ingresado:", n1) // ya tiene un espacio en blanco
		}
	}
}