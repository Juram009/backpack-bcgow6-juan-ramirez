package main
import "fmt"

func main(){
	age := 23
	employee := true
	antiquity := 2.0
	salary := 10000
	if age > 22 && employee && antiquity > 1 {
		if salary >=100000 {
			fmt.Println("Se aprobará el prestamo y no se cobrarán intereses")
		}else {
			fmt.Println("Se aprobará el prestamo, pero se cobrarán intereses")
		}
	} else {
		fmt.Println("No se aprueba el prestamo")
	}
}