package main
import (
	"fmt"
	"errors"
)

const (
	dog = "dog"
	cat = "cat"
	hamster = "hamster"
	tarantula = "tarantula"
)

func main() {
	animalDog, err := animal(dog)
	animalCat, err := animal(cat)
	animalHamster, err := animal(hamster)
	animalTarantula, err := animal(tarantula)
	if err == nil {
		var amount float64
		amount+=animalDog(5)
		amount+=animalCat(5)
		amount+=animalHamster(5)
		amount+=animalTarantula(5)
		fmt.Println(amount)
	}else {
		fmt.Println(err.Error())
	}
}

func animal(animal string) (function func(int) float64, err error) {
	switch animal {
	case dog:
		return dogFunc, nil
	case cat:
		return catFunc, nil
	case hamster:
		return hamsterFunc, nil
	case tarantula:
		return tarantulaFunc, nil
	default:
		err := errors.New("Animal invalido")
		return nil, err
	}
}

func dogFunc(dogs int) float64{
	return float64 (dogs)*10.0
}

func catFunc(cats int) float64{
	return float64 (cats)*5.0
}

func hamsterFunc(hamsters int) float64{
	return float64 (hamsters)*0.25
}

func tarantulaFunc(tarantulas int) float64{
	return float64 (tarantulas)*0.15
}