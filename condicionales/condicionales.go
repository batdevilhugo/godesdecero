package condicionales

import (
	"fmt"
	"runtime"
)

func SistemaOperatico() {
	os := runtime.GOOS
	if os == "Linux" {
		fmt.Println("Es linux")
	} else {
		fmt.Println("No es linux")
	}

	//& que pasa con las variables???
	switch os:= runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "windows":
		fmt.Println("Es windows")
	default:
		fmt.Printf("%s!\n", os); 
	}
}