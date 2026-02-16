package variables

import (
	"fmt"
	"time"
)

// variables globales con mayusculas se pueden utilizar en todo la funcion
var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func Enteras() {
	var x int
	x = 5
	y := int32(10)
	z := int64(100)

	fmt.Println(x, y, z)

	Nombre = "hugo"
	Estado = true
	Sueldo = 15000.46675
	Fecha = time.Now()
	fmt.Println(Nombre, Estado, Sueldo, Fecha)
}