package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id              string `json:"id"`
	Nombre          string `json:"nombre"`
	Color           string `json:"color"`
	Precio          int    `json:"precio"`
	Stock           int    `json:"stock"`
	Codigo          string `json:"codigo"`
	Publicado       bool   `json:"publicado"`
	FechaDeCreacion string `json:"fechaDeCreacion"`
}

func main() {
	router := gin.Default() //engine de gin
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{ //gin.H para parejas llave valor como en un map
			"message": "Hola Juan",
		})
	})
	router.GET("/productos", GetAll)
	router.GET("/productos/:Id", GetOne)
	router.Run()

	router = gin.Default()
}
func getProducts() ([]Producto, error) {
	var productos []Producto
	data, err := os.ReadFile("./products.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &productos)
	if err != nil {
		return nil, err
	}

	return productos, nil
}

func GetAll(ctx *gin.Context) {
	products, err := getProducts()
	var filtrados []*Producto
	for _, e := range products {
		if ctx.GetBool("Publicado") != e.Publicado {
			filtrados = append(filtrados, &e)
		}
	}
	if err != nil {
		ctx.String(400, "Error in JSON")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Productos": filtrados,
	})
}

func GetOne(ctx *gin.Context) {
	productos, err := getProducts()
	if err != nil {
		ctx.String(404, "Producto no encontrado")
		return
	}
	id := ctx.Param("Id")
	if err != nil {
		ctx.String(404, "Producto no encontrado")
		return
	}
	for _, value := range productos {
		if id == value.Id {
			ctx.JSON(200, value)
			return
		}
	}
	ctx.String(404, "Producto no encontrado")

}
func filter() {

}
