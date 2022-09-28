package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Producto struct {
	Nombre   string
	Precio   int
	Cantidad int
}

type Servicio struct {
	Nombre  string
	Precio  int
	Minutos int
}

type Mantenimiento struct {
	Nombre string
	Precio int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	canal := make(chan int)
	var productos []Producto
	for i := 1; i < 6; i++ {
		p := Producto{
			Nombre:   "Producto" + fmt.Sprint(i),
			Precio:   rand.Intn(300),
			Cantidad: rand.Intn(30),
		}
		productos = append(productos, p)
	}
	var servicios []Servicio
	for i := 1; i < 6; i++ {
		s := Servicio{
			Nombre:  "Servicio" + fmt.Sprint(i),
			Precio:  rand.Intn(300),
			Minutos: rand.Intn(300),
		}
		servicios = append(servicios, s)
	}
	var mantenimientos []Mantenimiento
	for i := 1; i < 6; i++ {
		m := Mantenimiento{
			Nombre: "Mantenimiento" + fmt.Sprint(i),
			Precio: rand.Intn(300),
		}
		mantenimientos = append(mantenimientos, m)
	}
	fmt.Printf("Productos: %v\n", productos)
	fmt.Printf("Servicios: %v\n", servicios)
	fmt.Printf("Mantenimientos: %v\n", mantenimientos)
	go costoProductos(productos, canal)
	go costoServicios(servicios, canal)
	go costoMantenimientos(mantenimientos, canal)
	lectura1 := <-canal
	lectura2 := <-canal
	lectura3 := <-canal
	fmt.Println("Monto final: ", lectura1+lectura2+lectura3)
}

func costoProductos(productos []Producto, ch chan int) {
	var res int
	for _, producto := range productos {
		res += producto.Cantidad * producto.Precio
	}
	ch <- res
}

func costoServicios(servicios []Servicio, ch chan int) {
	var res int
	for _, servicio := range servicios {
		if servicio.Minutos < 30 {
			res += servicio.Precio
		}
		res += (servicio.Minutos / 30) * servicio.Precio
	}
	ch <- res
}

func costoMantenimientos(mantenimientos []Mantenimiento, ch chan int) {
	var res int
	for _, mantenimiento := range mantenimientos {
		res += mantenimiento.Precio
	}
	ch <- res
}
