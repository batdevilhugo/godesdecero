package ejercicio

import "strconv"

func Ejercicio01(texto string) (bool, int) {
	numero, err := strconv.Atoi(texto)
	// una forma de manejar errores
	if err != nil {
		return false, 0
	}
	return true, numero
}