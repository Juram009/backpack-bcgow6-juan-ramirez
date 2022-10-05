package main

import (
	"fmt"
)

func main() {
	var salario int = 10000
	if salario < 150000 {
		error := fmt.Errorf("Error: El minimo imponible es de 150000 y el salario ingresado es de: $%d", salario)
		fmt.Println(error)
	} else {
		fmt.Println("Paga impuestos")
	}
}
