package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   int
	Cantidad int
}

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Producto
}

func main() {
	u := Usuario{
		Nombre:    "Pepe",
		Apellido:  "Gonzalez",
		Correo:    "pepegonzalez@gmail.com",
		Productos: []Producto{},
	}
	u.addProduct(newProduct("Papas", 1500), 15)
	fmt.Printf("%v\n", u)
	u.deleteProducts()
	fmt.Printf("%v\n", u)
}

func newProduct(nombre string, precio int) Producto {
	return Producto{
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: 0,
	}
}

func (u *Usuario) addProduct(p Producto, cantidad int) {
	p.Cantidad = cantidad
	u.Productos = append(u.Productos, p)
	fmt.Println("Producto agregado!")
}

func (u *Usuario) deleteProducts() {
	u.Productos = []Producto{}
}
