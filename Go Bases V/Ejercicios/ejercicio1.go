package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Ocurrió un error: %v \n", err)
		}
		fmt.Println("Ejecución finalizada")
	}()

	_, err := os.ReadFile("./customers.txt")
	if err != nil {
		panic("No se encontró el archivo")
	}
}
