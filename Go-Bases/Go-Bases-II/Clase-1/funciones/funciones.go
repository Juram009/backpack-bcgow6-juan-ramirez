package main
import "fmt"

const(
	suma = "+"
	resta = "-"
	multiplicacion = "*"
	division = "/"
)

func main() {
	a, b, c := 0, 2, -1
	inspeccionarVariable(a)
	inspeccionarVariable(b)
	inspeccionarVariable(c)
	//fmt.Println(suma(25.5, 12.4))
	fmt.Println(operacion(10, 5.5, suma))
	fmt.Println(operacion(10, 5.5, resta))
	fmt.Println(operacion(10, 5.5, multiplicacion))
	fmt.Println(operacion(10, 5.5, division))
}

func operacion(num1, num2 float64, operador string) float64{
	switch operador {
	case suma:
		return (num1+num2)
	case resta:
		return (num1-num2)
	case multiplicacion:
		return (num1*num2)
	case division:	
		return (num1/num2)
	}
	return 0
}

func inspeccionarVariable(numero int) {
	if numero > 0 {
		fmt.Println("Es positivo")
	} else if numero == 0 {
		fmt.Println("Es cero")
	} else {
		fmt.Println("Es negativo")
	}
}
/*
func suma(num1, num2 float64) float64{
	result := num1 + num2
	return result
}*/