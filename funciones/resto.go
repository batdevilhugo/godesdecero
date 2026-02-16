package funciones

import (
	"strconv"
)

func ConvierteTexto(numero int) (bool, string) { // devuelvo un bool y string
	var texto string
	texto = strconv.Itoa(numero)
	return true, texto // devolverlo en ese orden es buena practica
}

func ConvierteNumero(texto string) (bool, int) {
	numero, err := strconv.Atoi(texto)
	if err != nil {
		return false, 0
	}
	return true, numero
}
