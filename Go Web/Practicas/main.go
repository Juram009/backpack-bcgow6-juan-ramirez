package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id              string
	Nombre          string
	Color           string
	Precio          int
	Stock           int
	Codigo          string
	Publicado       bool
	FechaDeCreacion string
}

func main() {
	router := gin.Default() //engine de gin
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{ //gin.H para parejas llave valor como en un map
			"message": "Hola Juan",
		})
	})
	router.GET("/productos", GetAll)
	router.Run()

	router = gin.Default()
}
func generateProducts() []Producto {
	products := []Producto{
		{"1", "Televisor", "Negro", 1000000, 57, "df34jh174", true, "07/04/2020"},
		{"2", "Repisa", "Negro", 40000, 12, "jk3h4gf576", true, "13/08/2021"},
	}
	return products
}

func GetAll(ctx *gin.Context) {
	products := generateProducts()
	ctx.JSON(http.StatusOK, gin.H{
		"Productos": products,
	})
}
