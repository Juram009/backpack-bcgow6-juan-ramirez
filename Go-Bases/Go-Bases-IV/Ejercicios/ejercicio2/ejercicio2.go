package main

import (
	"errors"
	"fmt"
)

func main() {
	salario := 10000
	if salario < 150000 {
		fmt.Println(errors.New("Error: el salario es menor al mínimo"))
	} else {
		fmt.Println("Paga impuestos")
	}
}
