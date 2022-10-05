package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum int
	var res string
	fileContentAsBytes, err := os.ReadFile("../productos.csv")
	if err != nil {
		panic("Error leyendo el archivo")
	} else {
		for _, value := range fileContentAsBytes {
			if string(value) == ";" {
				res += fmt.Sprint("\t")
			} else {
				res += fmt.Sprintf(string(value))
			}
		}
		lines := strings.Split(res, "\n")
		for i := 1; i < len(lines)-1; i++ {
			product := strings.Split(lines[i], "\t")
			price, err := strconv.Atoi(product[1])
			if err != nil {
				panic("Archivo con errores")
			} else {
				sum += price
			}
		}
		fmt.Println(res)
		fmt.Println("TOTAL: $", sum)
	}
}
