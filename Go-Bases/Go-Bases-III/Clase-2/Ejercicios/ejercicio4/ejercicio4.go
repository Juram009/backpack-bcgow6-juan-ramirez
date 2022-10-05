package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canalInsercion := make(chan int64)
	canalBurbuja := make(chan int64)
	canalSeleccion := make(chan int64)
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)
	go insercion([][]int{variable1, variable2, variable3}, canalInsercion)
	go seleccion([][]int{variable1, variable2, variable3}, canalSeleccion)
	go burbuja([][]int{variable1, variable2, variable3}, canalBurbuja)
	fmt.Println("Inserción: ", <-canalInsercion, "ms")
	fmt.Println("Selección: ", <-canalSeleccion, "ms")
	fmt.Println("Burbuja: ", <-canalBurbuja, "ms")
}

func insercion(lista [][]int, ch chan int64) {
	now := time.Now()
	for k := 0; k < 3; k++ {
		var auxiliar int
		for i := 1; i < len(lista[k]); i++ {
			auxiliar = lista[k][i]
			for j := i - 1; j >= 0 && lista[k][j] > auxiliar; j-- {
				lista[k][j+1] = lista[k][j]
				lista[k][j] = auxiliar
			}
		}
	}
	ch <- time.Now().Sub(now).Milliseconds()
}

func seleccion(lista [][]int, ch chan int64) {
	now := time.Now()
	for k := 0; k < 3; k++ {
		for i := 0; i < len(lista[k]); i++ {
			min := i
			for j := i + 1; j < len(lista[k]); j++ {
				if lista[k][min] > lista[k][j] {
					min = j
				}
			}
			aux := lista[k][i]
			lista[k][i] = lista[k][min]
			lista[k][min] = aux
		}
	}
	ch <- time.Now().Sub(now).Milliseconds()
}

func burbuja(lista [][]int, ch chan int64) {
	now := time.Now()
	for k := 0; k < 3; k++ {
		var auxiliar int
		for i := 0; i < len(lista[k]); i++ {
			for j := 0; j < len(lista[k]); j++ {
				if lista[k][i] > lista[k][j] {
					auxiliar = lista[k][i]
					lista[k][i] = lista[k][j]
					lista[k][j] = auxiliar
				}
			}
		}
	}
	ch <- time.Now().Sub(now).Milliseconds()
}
