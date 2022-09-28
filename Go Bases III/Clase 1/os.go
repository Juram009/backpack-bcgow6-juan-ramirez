package main

import (
	"fmt"
	"os"
)

func main() {
	entries, err := os.ReadDir("..")
	if err != nil {
		panic("PANIC!")
	} else {
		fmt.Printf("Collected entries: %v\n", entries)
		for _, value := range entries {
			fmt.Println(value.Name())
		}
	}
	fileContentAsBytes, err := os.ReadFile("./fmt.go")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(fileContentAsBytes))
	e := os.WriteFile(
		"./fmt.txt",
		[]byte(string(fileContentAsBytes)),
		0644,
	)
	if e != nil {
		panic("Successfullyn't")
	} else {
		fmt.Println("Successfully!")
	}
}
func env() {

	err := os.Setenv("NAME", "Pepe")
	if err != nil {
		panic(err)
	}
	value := os.Getenv("NAME")
	fmt.Println(value)
	value, ok := os.LookupEnv("NAME")
	if !ok {
		panic("No se encuentra la variable de entorno")
	}
	fmt.Println(value)
}
func ex() {
	os.Exit(0)

}
