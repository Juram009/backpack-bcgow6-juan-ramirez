package main

import (
	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Web/Project/cmd/server/handler"
	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Web/Project/internal/products"
	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Web/Project/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	product := handler.NewProduct(service)

	route := gin.Default()

	pr := route.Group("/products")
	pr.POST("/", product.Store())
	pr.GET("/", product.GetAll())
	pr.GET("/:id", product.GetOne())
	pr.PUT("/:id", product.Update())
	pr.DELETE("/:id", product.Delete)
	pr.PATCH("/:id", product.UpdateNamePrice())
	if err := route.Run(); err != nil {
		panic(err)
	}
}
