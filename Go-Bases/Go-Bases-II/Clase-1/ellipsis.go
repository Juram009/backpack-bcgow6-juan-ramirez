package main
import "fmt"
const(
	opSuma = "+"
	opResta = "-"
	opMultiplicacion = "*"
	opDivision = "/"
)

func main() {
	fmt.Println(operacion(opSuma, 10, 5.5,6, 123, 123, -3256))
	fmt.Println(operacion(opResta, 10, 5.5))
	fmt.Println(operacion(opMultiplicacion, 10, 5.5))
	fmt.Println(operacion(opDivision, 10, 5.5))
}

func suma(valores []float64)float64 {
	var result float64
	for _, value := range valores {
		result += value
	}
	return result
}

func resta(valores []float64)float64 {
	var result float64
	for i, value := range valores {
		if i==0{
			result=value
		}else{
			result -= value
		}
	}
	return result
}

func multiplicacion(valores []float64)float64 {
	var result float64 =1
	for i, value := range valores {
		if i==0{
			result=value
		}else{
			result *= value
		}
	}
	return result
}

func division(valores []float64)float64 {
	var result float64 
	for i, value := range valores {
		if i==0{
			result=value
		}else{
			result /= value
		}
	}
	return result
}

func operacion(operador string, nums ... float64) float64{
	switch operador {
	case opSuma:
		return suma(nums)
	case opResta:
		return resta(nums)
	case opMultiplicacion:
		return multiplicacion(nums)
	case opDivision:	
		return division(nums)
	}
	return 0
}