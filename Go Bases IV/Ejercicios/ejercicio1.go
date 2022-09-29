package main

import (
	"fmt"
	"os"
)

type Error struct{}

func main() {
	salario := 10000
	error := &Error{}
	if salario < 150000 {
		fmt.Println("Error: ", error.Error())
		os.Exit(1)
	}
	fmt.Println("Paga impuestos")
}

func (e *Error) Error() string {
	return "Salario menor al minimo"
}
