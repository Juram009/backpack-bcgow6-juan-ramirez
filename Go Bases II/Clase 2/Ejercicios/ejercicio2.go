package main
import "fmt"

type Matrix struct {
	Alto int
	Ancho int
	Values []float64
	Cuadratica bool
	Max int
}

func (m *Matrix) set(values ...float64){
	m.Values = values
}

func (m Matrix) print(){
	for k := 0;k < len(m.Values);k++ {
		//i := k / m.Ancho
		j := k % m.Ancho
		if j==0 {
			fmt.Printf("\n")
		}
		fmt.Print(m.Values[k]," ")
	}
	fmt.Printf("\n")
}

func main() {
	m := Matrix{
		Ancho: 6,
		Alto: 2,
		Cuadratica: false,
		Max: 65,
	}

	m.set(4, 5, 7, 2, 7, 9, 5, 7, 2, 10, 65, 23)
	m.print()
}