package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Testing/Go-Testing-II/clase-2/ejercicios/Project/internal/products"
	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Testing/Go-Testing-II/clase-2/ejercicios/Project/pkg/web"
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

// ListProducts godoc
// @Summary  Show list products
// @Tags     Products
// @Produce  json
// @Param    token  header    string        true  "token"
// @Success  200    {object}  web.Response  "List products"
// @Failure  401    {object}  web.Response  "Unauthorized"
// @Failure  500    {object}  web.Response  "Internal server error "
// @Failure  404    {object}  web.Response  "Not found products"
// @Router   /products [GET]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Invalid Token"))
			return
		}

		product, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}
		if len(product) == 0 {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "There are no products"))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, product, ""))
	}
}

// GetProduct godoc
// @Summary  Show one products by id
// @Tags     Products
// @Produce  json
// @Param    id     path      int            true  "Id product"
// @Param    token  header    string         true  "token"
// @Success  200    {object}  web.Response  "Product"
// @Failure  401    {object}  web.Response  "Unauthorized"
// @Failure  404    {object}  web.Response  "Product Not found"
// @Router   /products/{id} [GET]
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
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, product, ""))
	}
}

// Store Product godoc
// @Summary  Store product
// @Tags     Products
// @Accept   json
// @Produce  json
// @Param    token    header    string          true  "token requerido"
// @Param    product  body      Product  true  "Product to Store"
// @Success  201      {object}  web.Response "Product"
// @Failure  401      {object}  web.Response "Unauthorized"
// @Failure  400      {object}  web.Response "Bad Request"
// @Failure  404      {object}  web.Response "Product Not Found"
// @Router   /products [POST]
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
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, product, ""))
	}
}

// UpdateProduct godoc
// @Summary  Update product
// @Tags     Products
// @Accept   json
// @Produce  json
// @Param    id       path      int             true   "Id product"
// @Param    token    header    string          false  "Token"
// @Param    product  body      Product  true   "Product to update"
// @Success  200      {object}  web.Response "Product"
// @Failure  401      {object}  web.Response "Unauthorized"
// @Failure  400      {object}  web.Response "Bad Request"
// @Failure  404      {object}  web.Response "Product Not Found"
// @Router   /products/{id} [PUT]
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
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, product, ""))
	}
}

// Delete Product
// @Summary  Delete product
// @Tags     Products
// @Param    id     path      int     true  "Product id"
// @Param    token  header    string  true  "Token"
// @Success  202
// @Failure  401      {object}  web.Response "Unauthorized"
// @Failure  400      {object}  web.Response "Bad Request"
// @Failure  404      {object}  web.Response "Product Not Found"
// @Router   /products/{id} [DELETE]
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

// Update Name Price Product godoc
// @Summary      Update name and price product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Description  This endpoint update field name and price from a product
// @Param        token  header    string               true  "Token header"
// @Param        id     path      int                  true  "Product Id"
// @Param        name   body      ProductRequestPatch  true  "Product name"
// @Success  200      {object}  web.Response "Product"
// @Failure  401      {object}  web.Response "Unauthorized"
// @Failure  400      {object}  web.Response "Bad Request"
// @Failure  404      {object}  web.Response "Product Not Found"
// @Router       /products/{id} [PATCH]
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

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, product, ""))
	}
}
