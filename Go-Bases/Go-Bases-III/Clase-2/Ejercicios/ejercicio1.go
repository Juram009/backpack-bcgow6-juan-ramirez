package main

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func cambiarStr(puntero *string, value string) {
	*puntero = value
}

func cambiarEdad(puntero *int, value int) {
	*puntero = value
}

func main() {
	u := Usuario{
		Nombre:     "Pepe",
		Apellido:   "Gonzalez",
		Edad:       32,
		Correo:     "pepegonzalez@gmail.com",
		Contraseña: "PepeGonzalez123",
	}
	fmt.Printf("%v\n", u)
	cambiarStr(&u.Nombre, "Joaquin")
	cambiarEdad(&u.Edad, 34)
	fmt.Printf("%v\n", u)
}
