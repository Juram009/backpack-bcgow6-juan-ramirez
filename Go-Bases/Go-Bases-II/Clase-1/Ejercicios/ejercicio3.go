package main
import "fmt"

func main() {
	salary := (calculateSalary(6000.0, "C"))
	fmt.Println("El salario del empleado es de",salary)
}

func calculateSalary(minutes float64, category string) float64 {
	switch category {
	case "A":
		basicSalary := (minutes/60)*3000
		aditional := basicSalary*0.5
		return basicSalary+aditional
	case "B":
		basicSalary := (minutes/60)*1500
		aditional := basicSalary*0.2
		return basicSalary+aditional
	case "C":	
		basicSalary := (minutes/60)*1000
		return basicSalary
	}
	return 0
}