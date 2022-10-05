package main
import "fmt"

func main() {
	salario := 150000.0
	imp := impuesto(salario)
	fmt.Println("El impuesto es del", imp, "% del salario, es decir de", salario*(imp/100.0))
}

func impuesto(salario float64) float64{
	if salario >= 150000 {
		return 10
	}else if salario >= 50000 {
		return 17
	}
	return 0
}