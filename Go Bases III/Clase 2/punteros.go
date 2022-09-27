package main
import "fmt"

func main() {
	//var puntero *int
	//p2 := new(int)

	var numero int
	p3 := &numero
	fmt.Printf("Valor: %d\n", numero)
	Incrementar(p3)
	fmt.Printf("Valor: %d\n", numero)
	fmt.Printf("La direccion es: %p\n", p3)
	fmt.Printf("La direccion es: %p\n", &numero)
}

func Incrementar(puntero *int){
	*puntero = 20 
}
