package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	//Arrange
	p1 := Product{
		Id:           1,
		Name:         "MacBook",
		Color:        "Gris",
		Price:        100000,
		Stock:        7,
		Code:         "8403",
		Published:    false,
		CreationDate: "29/03/2019",
	}
	p2 := Product{
		Id:           2,
		Name:         "Pantalla",
		Color:        "Negra",
		Price:        1000,
		Stock:        2,
		Code:         "2639",
		Published:    true,
		CreationDate: "09/03/2020",
	}
	products := []Product{p1, p2}
	expectedProducts := []Product{p1, {
		Id:           2,
		Name:         "Pantall",
		Color:        "Negra",
		Price:        1000,
		Stock:        2,
		Code:         "2639",
		Published:    true,
		CreationDate: "09/03/2020",
	}}
	mockStorage := MockStorage{
		dataMock: products,
	}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	p3 := Product{
		Id:           2,
		Name:         "Pantall",
		Color:        "Negra",
		Price:        1000,
		Stock:        2,
		Code:         "2639",
		Published:    true,
		CreationDate: "09/03/2020",
	}
	//Act
	p2, errPut := service.Update(p3.Id, p3.Name, p3.Color, p3.Price, p3.Stock, p3.Code, p3.Published, p3.CreationDate)
	res, errGet := service.GetAll()

	//Assert
	assert.Equal(t, expectedProducts, res)
	assert.Nil(t, errPut)
	assert.Nil(t, errGet)
	assert.True(t, mockStorage.ReadWasCalled)
}

func TestDelete(t *testing.T) {
	//Arrange
	p1 := Product{
		Id:           1,
		Name:         "MacBook",
		Color:        "Gris",
		Price:        100000,
		Stock:        7,
		Code:         "8403",
		Published:    false,
		CreationDate: "29/03/2019",
	}
	p2 := Product{
		Id:           2,
		Name:         "Pantalla",
		Color:        "Negra",
		Price:        1000,
		Stock:        2,
		Code:         "2639",
		Published:    true,
		CreationDate: "09/03/2020",
	}
	products := []Product{p1, p2}
	expectedProducts := []Product{p1}
	mockStorage := MockStorage{
		dataMock: products,
	}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	//Act
	errDelExistent := service.Delete(2)
	errDelInexistent := service.Delete(4)
	res, errGet := service.GetAll()

	//Assert
	assert.Nil(t, errDelExistent)
	assert.Nil(t, errGet)
	assert.Equal(t, expectedProducts, res)
	assert.NotNil(t, errDelInexistent)
	assert.True(t, mockStorage.ReadWasCalled)
}
