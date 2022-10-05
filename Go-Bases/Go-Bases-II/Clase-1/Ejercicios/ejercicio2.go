package main
import "fmt"

func main() {
	prom := promedio(4,3,5,2,4,2,3,4)
	fmt.Println("El promedio del estudiante es", prom)
}

func promedio(notas ...float64)float64{
	var resultado float64
	for _, valor := range notas {
		resultado += valor
	}
	return resultado/float64 (len(notas))
}