package main

import (
	"fmt"
	"math/rand"
	"os"
)

type Cliente struct {
	Legajo    int
	DNI       int
	Nombre    string
	Telefono  string
	Domicilio string
}

func main() {
	cliente := Cliente{
		DNI:       123,
		Nombre:    "Francisco Perez",
		Telefono:  "3192740283",
		Domicilio: "Calle 3 #24-65",
	}
	err := cliente.generarID()
	if err != nil {
		panic(err)
	}
	verificarCliente(cliente.Legajo)
	NuevoCliente(2134, 35457, "Jorge Lopez", "3104857234", "Carrera 50 #56-21")
	fmt.Println("Fin de la ejecución")
}

func (c *Cliente) generarID() error {
	randomID := rand.Int()
	if randomID == 0 {
		return fmt.Errorf("El id no puede ser 0")
	} else {
		c.Legajo = randomID
		return nil
	}
}

func verificarCliente(id int) {
	_, err := os.Open("./customer.txt")
	if err != nil {
		panic("Error: el archivo indicado no fue encontrado o está dañado")
	}
}

func NuevoCliente(legajo, dni int, nombre, telefono, domicilio string) (*Cliente, error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error")
		}
	}()

	return &Cliente{
		Legajo:    legajo,
		DNI:       dni,
		Nombre:    nombre,
		Telefono:  telefono,
		Domicilio: domicilio,
	}, nil
}
