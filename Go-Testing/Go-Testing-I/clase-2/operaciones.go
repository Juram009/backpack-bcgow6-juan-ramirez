package main

import "fmt"

func main() {
	fmt.Println(Add(10, 5))
}

func Add(a, b int) (int, error) {
	return a + b, nil
}

func Sub(a, b int) int {
	return a - b
}
