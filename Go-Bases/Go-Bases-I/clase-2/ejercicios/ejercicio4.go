package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {
	fmt.Println("La edad de Benjamin es:", employees["Benjamin"], "años")
	ages21()
	addFede()
	deletePedro()
}

func ages21() {
	ages := 0
	for key := range employees {
		if employees[key] > 21 {
			ages++
		}
	}
	fmt.Println("Hay", ages, "empleados mayores de 21 años")
}

func addFede() {
	employees["Federico"] = 25
	fmt.Println(employees)
}
func deletePedro() {
	delete(employees, "Pedro")
	fmt.Println(employees)
}
