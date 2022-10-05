package main
import (
	"fmt"
	"errors"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)
 
func main() {
	minFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)
	if err == nil {
		fmt.Println("minValue:",minFunc(2, 3, 3, 4, 10, 2, 4, 5))
		fmt.Println("averageValue:",averageFunc(2, 3, 3, 4, 1, 2, 4, 5))
		fmt.Println("maxValue:",maxFunc(2, 3, 3, 4, 1, 2, 4, 5))
	}else {
		fmt.Println(err.Error())
	}
}

func operation (operador string) (func(valores ...float64) float64,error) {
	switch operador {
	case minimum:
		return opMinimum, nil
	case average:
		return opAverage, nil
	case maximum:
		return opMaximum, nil
	}
	return nil, errors.New("Operación invalida")
}

func opMinimum(valores ...float64) float64{
	var min float64
	for i,value := range valores {
		if i==0 {
			min = value
		} else {
			if value < min {
				min = value
			}
		}
	}
	return min
}

func opAverage(valores ...float64) float64{
	sum := 0.0
	for _, value := range valores {
		sum += value
	}
	return sum/float64 (len(valores)) 
}

func opMaximum(valores ...float64) float64{
	var max float64
	for i,value := range valores {
		if i==0 {
			max = value
		} else {
			if value > max {
				max = value
			}
		}
	}
	return max
}