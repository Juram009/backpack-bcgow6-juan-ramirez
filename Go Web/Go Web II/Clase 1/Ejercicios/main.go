package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id              int    `json:"id" binding:"required"`
	Nombre          string `json:"nombre" binding:"required"`
	Color           string `json:"color" binding:"required"`
	Precio          int    `json:"precio" binding:"required"`
	Stock           int    `json:"stock" binding:"required"`
	Codigo          string `json:"codigo" binding:"required"`
	Publicado       bool   `json:"publicado" binding:"required"`
	FechaDeCreacion string `json:"fechaDeCreacion" binding:"required"`
}

var productos []Producto
var lastID int

const token string = "123"

func main() {
	router := gin.Default()
	p := router.Group("products")
	p.POST("", Save())
	router.Run()
}

func Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if token != ctx.GetHeader("token") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Token invalido"})
		}
		var req Producto
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		if req.Codigo == "" || req.Color == "" || req.FechaDeCreacion == "" || req.Nombre == "" || req.Precio == 0 || req.Stock == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Todos los campos son requeridos"})
			return
		}
		lastID++
		req.Id = lastID
		productos = append(productos, req)
		ctx.JSON(http.StatusOK, req)
	}
}
