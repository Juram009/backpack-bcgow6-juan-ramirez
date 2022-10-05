package main
import (
	"fmt"
	"time"
)

func procesar(i int, ch chan int){
	fmt.Println(i, "- Inicio")
	time.Sleep(1000*time.Millisecond)
	fmt.Println(i, "- Final")
	ch <- i 
}

func main(){
	canal := make(chan int)
	now := time.Now()
	fmt.Println("Inicio de programa")
	for i:=0 ;i < 10; i++{
		go procesar(i,canal)
	}
	for i:=0 ;i < 10; i++{
		lectura := <- canal
		fmt.Println("Lectura de canal:",lectura)
	}
	
	//time.Sleep(1000*time.Millisecond)
	fmt.Println("Fin de programa, ",time.Now().Sub(now))
}