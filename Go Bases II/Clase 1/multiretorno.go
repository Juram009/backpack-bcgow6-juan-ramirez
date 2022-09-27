package main
import (
	"fmt"
	"errors"
)

func main() {
	s, r, m, d := calcular(10,2)
	fmt.Println("Suma:", s)
	fmt.Println("Resta:", r)
	fmt.Println("Multiplicación:", m)
	fmt.Println("División:", d)
	fmt.Println(calcular(12,45))
	result, err := division(10,0)
	if err!= nil {
		panic(err)
	}
	fmt.Println(result)
}

func calcular (num1, num2 float64) (float64, float64, float64, float64) {
	suma := num1 + num2
	resta := num1 - num2
	multiplicacion := num1 * num2
	division := num1 / num2
	return suma, resta, multiplicacion, division
}

func division(num1, num2 float64) (float64, error) {
	if num2 ==0 {
		return 0, errors.New("El divisor no puede ser cero")
	}
	return num1 / num2, nil
}