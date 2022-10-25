package main

import (
	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Testing/desafio-cierre-testing-main/cmd/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	r.Run(":18085")

}
