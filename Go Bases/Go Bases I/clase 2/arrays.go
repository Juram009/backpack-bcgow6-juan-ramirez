package main
import "fmt"
func main(){
	var myOtherArray = [...]int{1, -2, 3, 4, 34, 123,7}
    fmt.Println("myOtherArray", myOtherArray)
    fmt.Println("//------- Slice ----") //Tiene un array como base. El tamanio puede cambiar dinamicamente
    slice := myOtherArray[2:5]
    fmt.Println("slice", slice)
    fmt.Println("slice len:", len(slice))
    fmt.Println("cap len:", cap(slice))
}