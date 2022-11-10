package main

import (
	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Implementacion-Storage/Project/cmd/server/routes"
	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Implementacion-Storage/Project/pkg/db"
)

func main() {
	engine, db := db.ConnectDatabase()
	router := routes.NewRouter(engine, db)
	router.MapRoutes()

	engine.Run(":8080")
}
