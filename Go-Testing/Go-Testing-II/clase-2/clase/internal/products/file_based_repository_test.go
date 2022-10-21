package products

import (
	"testing"

	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Testing/Go-Testing-II/clase-2/clase/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	// Arrange.
	database := []domain.Product{
		{
			ID:    1,
			Name:  "Caja de galletitas Boreo 1kg",
			Type:  "Galletitas y snacks",
			Count: 2000,
			Price: 300,
		},
		{
			ID:    2,
			Name:  "Desodorante Roxana",
			Type:  "Higiene personal",
			Count: 5000,
			Price: 120,
		},
	}

	mockStorage := MockStorage{
		dataMock: database,
	}

	repository := NewFileBasedRepository(&mockStorage)

	// Act.
	result, err := repository.GetAll()

	// Assert.
	assert.Nil(t, err)
	assert.Equal(t, database, result)
}

func TestStore(t *testing.T) {
	// Arrange.
	expected := []domain.Product{
		{
			ID:    1,
			Name:  "Caja de galletitas Boreo 1kg",
			Type:  "Galletitas y snacks",
			Count: 2000,
			Price: 300,
		},
		{
			ID:    2,
			Name:  "Desodorante Roxana",
			Type:  "Higiene personal",
			Count: 5000,
			Price: 120,
		},
		{
			ID:    3,
			Name:  "Cola Cola 500ml",
			Type:  "Gaseosas",
			Count: 12000,
			Price: 172,
		},
	}

	initialDatabase := []domain.Product{
		{
			ID:    1,
			Name:  "Caja de galletitas Boreo 1kg",
			Type:  "Galletitas y snacks",
			Count: 2000,
			Price: 300,
		},
		{
			ID:    2,
			Name:  "Desodorante Roxana",
			Type:  "Higiene personal",
			Count: 5000,
			Price: 120,
		},
	}

	mockStorage := MockStorage{
		dataMock: initialDatabase,
	}

	repository := NewFileBasedRepository(&mockStorage)

	// Act.
	productToCreate := domain.Product{
		ID:    3,
		Name:  "Cola Cola 500ml",
		Type:  "Gaseosas",
		Count: 12000,
		Price: 172,
	}

	result, err := repository.Store(productToCreate)

	// Assert.
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, expected)
	assert.Equal(t, productToCreate, result)
}
