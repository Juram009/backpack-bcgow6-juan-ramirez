package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllBySeller(t *testing.T) {
	//Arrange
	expected := []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}
	repository := NewRepository()

	//Act
	res, err := repository.GetAllBySeller("Seller01")

	//Assert
	assert.Equal(t, expected, res)
	assert.Nil(t, err)
}
