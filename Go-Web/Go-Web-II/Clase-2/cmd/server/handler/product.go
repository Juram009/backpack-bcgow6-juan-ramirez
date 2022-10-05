package handler

import (
	"net/http"
	"strconv"

	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Web/Go-Web-II/Clase-2/internal/products"
	"github.com/gin-gonic/gin"
)

type request struct {
	Id           int    `json:"id" binding:"required"`
	Name         string `json:"nombre" binding:"required"`
	Color        string `json:"color" binding:"required"`
	Price        int    `json:"precio" binding:"required"`
	Stock        int    `json:"stock" binding:"required"`
	Code         string `json:"codigo" binding:"required"`
	Published    bool   `json:"publicado" binding:"required"`
	CreationDate string `json:"fecha" binding:"required"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "12345" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Invalid Token",
			})
			return
		}

		product, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusAccepted, product)
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "12345" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Token"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		product, err := p.service.Store(req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, product)
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("Id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Id"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		if req.Code == "" || req.Color == "" || req.CreationDate == "" || req.Name == "" || req.Price == 0 || req.Stock == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "All fields are required"})
			return
		}
		product, err := p.service.Update(int(id), req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, product)
	}
}
