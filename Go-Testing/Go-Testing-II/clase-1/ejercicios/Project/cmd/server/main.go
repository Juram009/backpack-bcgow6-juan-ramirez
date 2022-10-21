package main

import (
	"os"

	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Testing/Go-Testing-II/clase-1/ejercicios/Project/cmd/server/handler"
	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Testing/Go-Testing-II/clase-1/ejercicios/Project/docs"
	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Testing/Go-Testing-II/clase-1/ejercicios/Project/internal/products"
	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Testing/Go-Testing-II/clase-1/ejercicios/Project/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Bootcamp Go - API GO-WEB
// @version         1.0
// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	product := handler.NewProduct(service)

	route := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	api := route.Group("/api/v1")
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := api.Group("/products")
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
