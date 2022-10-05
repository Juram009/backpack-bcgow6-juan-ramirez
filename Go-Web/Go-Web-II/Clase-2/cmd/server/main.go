package main

import (
	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Web/Go-Web-II/Clase-2/cmd/server/handler"
	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Web/Go-Web-II/Clase-2/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)

	product := handler.NewProduct(service)

	route := gin.Default()

	pr := route.Group("/products")
	pr.POST("/", product.Store())
	pr.GET("/", product.GetAll())
	pr.PUT("/:id", product.Update())
	route.Run()
}
