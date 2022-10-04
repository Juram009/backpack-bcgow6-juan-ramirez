package main
import "fmt"

type Estudiante struct {
	Nombre 		string
	Apellido 	string
	DNI 		int
	Fecha 		int
}

func (e Estudiante) Detalle() {
	fmt.Println("Nombre:",e.Nombre)
	fmt.Println("Apellido:",e.Apellido)
	fmt.Println("DNI:",e.DNI)
	fmt.Println("Fecha:",e.Fecha)
}
func main() {
	student := Estudiante{
		Nombre: "Pepe",
		Apellido: "Suarez",
		DNI: 139849023,
		Fecha: 2022,
	}
	student.Detalle()
}
