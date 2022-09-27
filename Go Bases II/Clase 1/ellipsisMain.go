package main
import "fmt"

const(
	suma = "+"
	resta = "-"
	multiplicacion = "*"
	division = "/"
)

func main() {
	fmt.Println(operacion(suma, 10, 5.5,6, 123, 123, -3256))
	fmt.Println(operacion(resta, 10, 5.5))
	fmt.Println(operacion(multiplicacion, 10, 5.5))
	fmt.Println(operacion(division, 10, 5.5))
}

func opSuma(a, b float64) float64{
	return a+b
}
func opResta(a, b float64) float64{
	return a-b
}
func opMultiplicacion(a, b float64) float64{
	return a*b
}
func opDivision(a, b float64) float64{
	if b==0{
		return 0
	}
	return a/b
}

func orqOperacion(valores []float64, operacion func(num1, num2 float64) float64) float64 {
	var resultado float64
	for i, valor := range valores {
		if i == 0 {
			resultado = valor
		} else {
			resultado = operacion(resultado, valor)
		}
	}
	return resultado
}

func operacion(operador string, valores ...float64) float64 {
	switch operador {
	case suma:
		return orqOperacion(valores, opSuma)
	case resta:
		return orqOperacion(valores, opResta)
	case multiplicacion:
		return orqOperacion(valores, opMultiplicacion)
	case division:
		return orqOperacion(valores, opDivision)
	}
	return 0
}