package main

import (
	"fmt"
	"math/rand"
	"os"
)

type Producto struct {
	Id       int
	Precio   int
	Cantidad int
}

func main() {
	var productos []Producto
	for i := 0; i < 10; i++ {
		p := Producto{
			Id:       rand.Intn(9999),
			Precio:   rand.Intn(300),
			Cantidad: rand.Intn(30),
		}
		productos = append(productos, p)
	}
	txt := "Id;Precio;Cantidad\n"
	for i := 0; i < len(productos); i++ {
		txt += fmt.Sprintf("%d;%d;%d\n", productos[i].Id, productos[i].Precio, productos[i].Cantidad)
		e := os.WriteFile(
			"./productos.csv",
			[]byte(txt),
			0644,
		)
		if e != nil {
			panic("Successfullyn't")
		}
	}
	fmt.Println("Successfully!")
}
