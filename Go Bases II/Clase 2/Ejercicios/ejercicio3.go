package main
import "fmt"

type tienda struct {
	Productos []producto
}

type producto struct {
	Nombre string
	Tipo string
	Precio float64
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(producto)
}

func main() {
	tienda := nuevaTienda()
	p1 := nuevoProducto("Papas", "pequeño", 120)
	p2 := nuevoProducto("Escritorio", "mediano", 800)
	p3 := nuevoProducto("Moto", "grande", 3000)

	tienda.Agregar(p1)
	tienda.Agregar(p2)
	tienda.Agregar(p3)
	
	fmt.Println(tienda.Total())
}

func nuevoProducto(nombre, tipo string, precio float64) producto{
	p := producto{
		Nombre: nombre,
		Tipo: tipo,
		Precio: precio,
	}
	return p
}

func nuevaTienda() tienda{
	return tienda{}
}

func (t tienda) Total() float64{
	var total float64
	for _,p := range t.Productos {
		total += p.CalcularCosto()
	}
	return total
}

func (t *tienda) Agregar(p producto) {
	t.Productos = append(t.Productos, p)
}

func (p producto) CalcularCosto() float64{
	aditional := 0.0
	switch p.Tipo {
	case "mediano":
		aditional = p.Precio*.03
	case "grande":
		aditional = p.Precio*0.06 +2500
	}
	return p.Precio+aditional
}