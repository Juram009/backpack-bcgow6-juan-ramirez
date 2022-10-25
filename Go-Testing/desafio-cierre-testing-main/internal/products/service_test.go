package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessOnGetAllBySeller(t *testing.T) {
	//Arrange
	db := []Product{
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
	expected := []Product{
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
	}
	mockDB := MockDB{
		db: db,
	}
	service := NewService(&mockDB)
	//Act

	res, err := service.GetAllBySeller("Seller01")

	//Assert
	assert.Equal(t, expected, res)
	assert.Nil(t, err)
}
func TestErrorOnGetAllBySeller(t *testing.T) {
	//Arrange
	db := []Product{}
	mockDB := MockDB{
		db: db,
	}
	service := NewService(&mockDB)
	//Act

	res, err := service.GetAllBySeller("Seller01")

	//Assert
	assert.Equal(t, []Product(nil), res)
	assert.NotNil(t, err)
}
