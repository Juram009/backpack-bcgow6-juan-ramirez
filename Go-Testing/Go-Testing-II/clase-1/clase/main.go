package main

import (
	"fmt"

	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Testing/Go-Testing-II/clase-1/clase/directorio"
)

func main() {
	db := directorio.NewDB()
	motor := directorio.NewEngine(db)

	fmt.Printf("La versión actual es %s\n", motor.GetVersion())
}
