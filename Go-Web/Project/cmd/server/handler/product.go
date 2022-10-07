package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Web/Project/internal/products"
	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Web/Project/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Id           int    `json:"id"`
	Name         string `json:"name" binding:"required"`
	Color        string `json:"color" binding:"required"`
	Price        int    `json:"price" binding:"required"`
	Stock        int    `json:"stock" binding:"required"`
	Code         string `json:"code" binding:"required"`
	Published    bool   `json:"published" binding:"required"`
	CreationDate string `json:"date" binding:"required"`
}

type Product struct {
	service products.Service
}

type ProductRequestPatch struct {
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Invalid Token"))
			return
		}

		product, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusAccepted, web.NewResponse(http.StatusAccepted, product, ""))
	}
}

func (p *Product) GetOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Invalid Token"))
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		product, err := p.service.GetOne(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusAccepted, web.NewResponse(http.StatusAccepted, product, ""))
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Invalid Token"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		product, err := p.service.Store(req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusAccepted, web.NewResponse(http.StatusAccepted, product, ""))
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Invalid Token"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Invalid Id"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		if req.Code == "" || req.Color == "" || req.CreationDate == "" || req.Name == "" || req.Price == 0 || req.Stock == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "All fields are required"))
			return
		}
		product, err := p.service.Update(int(id), req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusAccepted, web.NewResponse(http.StatusAccepted, product, ""))
	}
}

func (p *Product) Delete(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Invalid Token"))
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Invalid Id "+err.Error()))
		return
	}

	if err := p.service.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusAccepted, web.NewResponse(http.StatusAccepted, "Successful Deleted", ""))
}

func (p *Product) UpdateNamePrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Invalid Token"))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Invalid Id "+err.Error()))
			return
		}

		var req ProductRequestPatch
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		product, err := p.service.UpdateNamePrice(id, req.Name, req.Price)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusAccepted, web.NewResponse(http.StatusAccepted, product, ""))
	}
}
