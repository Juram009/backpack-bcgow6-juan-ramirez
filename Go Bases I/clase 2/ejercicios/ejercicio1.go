package main
import "fmt"

func main(){
	palabra := "palabra"
	fmt.Println(len(palabra))
	for i := 0; i < len(palabra); i++{
		fmt.Printf("%c ", palabra[i])//%c es para el caracter porque palabra[i] devuelve el codigo ascii de la letra
	}
	print("\n")
}