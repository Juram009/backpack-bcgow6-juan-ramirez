package products

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	db := GetProducts()
	MockDB := MockDB{
		db: db,
	}
	service := NewService(&MockDB)
	handler := NewHandler(service)
	r := gin.Default()
	r.GET("/products", handler.GetProducts)
	return r
}

func createBadServer() *gin.Engine {
	MockDB := MockDB{
		db: []Product{},
	}
	service := NewService(&MockDB)
	handler := NewHandler(service)
	r := gin.Default()
	r.GET("/products", handler.GetProducts)
	return r
}

func createRequest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func TestGetProducts(t *testing.T) {
	//Arrange
	r := createServer()
	req, res := createRequest(http.MethodGet, "/products?seller_id=Seller01", "")
	//Act
	r.ServeHTTP(res, req)
	//Assert
	assert.Equal(t, http.StatusOK, res.Code)
}
func TestErr400GetProducts(t *testing.T) {
	//Arrange
	r := createServer()
	req, res := createRequest(http.MethodGet, "/products?seller_id=", "")
	//Act
	r.ServeHTTP(res, req)
	//Assert
	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func TestErr500GetProducts(t *testing.T) {
	//Arrange
	r := createBadServer()
	req, res := createRequest(http.MethodGet, "/products?seller_id=Seller01", "")
	//Act
	r.ServeHTTP(res, req)
	//Assert
	assert.Equal(t, http.StatusInternalServerError, res.Code)
}

func GetProducts() []Product {
	return []Product{
		{
			ID:          "123",
			SellerID:    "Seller01",
			Description: "Description of product",
			Price:       54.6,
		},
		{
			ID:          "456",
			SellerID:    "Seller01",
			Description: "Description of product",
			Price:       23.9,
		},
		{
			ID:          "789",
			SellerID:    "Seller02",
			Description: "Description of product",
			Price:       4.6,
		},
	}
}
