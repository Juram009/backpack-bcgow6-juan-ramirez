package handler

import (
	"net/http"
	"strconv"

	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Implementacion-Storage/Project/internal/domain"
	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Implementacion-Storage/Project/internal/product"
	"github.com/gin-gonic/gin"
)

type Product struct {
	service product.Service
}

func NewProduct(service product.Service) *Product {
	return &Product{
		service: service,
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.service.GetAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, products)
	}
}

func (p *Product) GetOne() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id, err := strconv.Atoi((ctx.Param("id")))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		product, err := p.service.Get(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, product)
	}
}

func (p *Product) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var product domain.Product
		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		product, err = p.service.Save(ctx, product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"product": product.Name + " added"})
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt((ctx.Param("id")), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = p.service.Delete(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusNoContent, gin.H{"delete": id})
	}
}
