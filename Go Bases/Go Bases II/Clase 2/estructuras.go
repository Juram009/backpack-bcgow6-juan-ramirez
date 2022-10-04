package main
import (
	"fmt"
	"encoding/json"

)

type Perro struct {
	Nombre string `json:"nombre"`
	Edad int `json:"edad"`
	Raza string `json:"raza"`
	Peso float64 `json:"peso"`
	Color string `json:"color"`
	Owner Dueño `json:"owner"`
}

func (p Perro) Ladrar(){
	fmt.Println("Guau!")
}

func (p Perro) Dormir() {
	fmt.Println("zzz")
}

func (p Perro) Comer() {
	fmt.Println("Ñom ñom ñom")
}

func (p *Perro) Renombrar(new string){ //Se usa puntero para hacer cambios en la estructura
	p.Nombre = new
	fmt.Println("Ahora me llamo",p.Nombre)
}

type Dueño struct{
	Documento int
	Nombre string
	Telefono int
}

type Gato struct {
	Nombre string `json:"nombre"`
	Edad int `json:"edad"`
	Raza string `json:"raza"`
	Peso float64 `json:"peso"`
	Caracter string `json:"caracter"`
}

func (g Gato) Dormir() {
	fmt.Println("zzz")
}

type Animal interface {
	Dormir()
	Comer()
}

func NewPerro() Animal {
	return &Perro{}
}

type ListaHeterogenea struct{
	Datos []interface{}
}

func main() {
	d1 := Dueño {
		Documento: 1234,
		Nombre: "Don Pepe",
		Telefono: 3102345678,
	}
	dog := Perro{
		Nombre:"Pepe", 
		Edad: 2, 
		Raza: "Golden", 
		Peso: 10,
		Owner: d1,
	}

	fmt.Printf("Dog: %+v\n", dog)
	fmt.Printf("Doc: %+v\n", dog.Owner.Documento)
	b, err := json.Marshal(&dog)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nsjon.Marshal: %+v\n",string (b))
	dog.Ladrar()
	dog.Renombrar("Yepeto")
	fmt.Println(dog.Nombre)

	a := NewPerro()
	a.Comer()
	a.Dormir()

	lista :=ListaHeterogenea{}
	lista.Datos = append(lista.Datos, 1)
	lista.Datos = append(lista.Datos, 5.6)
	lista.Datos = append(lista.Datos, "Hola")
	fmt.Println(lista.Datos)

	e, ok := lista.Datos[0].(int)
	fmt.Println(ok)
	fmt.Println(e)
}