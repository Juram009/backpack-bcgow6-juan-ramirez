package products

import (
	"errors"
	"testing"

	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Testing/Go-Testing-II/clase-2/clase/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationGetAll(t *testing.T) {
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
	service := NewService(repository)

	// Act.
	results, err := service.GetAll()

	// Assert.
	assert.Nil(t, err)
	assert.Equal(t, database, results)
}

func TestServiceIntegrationGetAllFail(t *testing.T) {
	// Arrange.
	expectedErr := errors.New("hello, i'm an error :D")

	mockStorage := MockStorage{
		dataMock:   nil,
		errOnWrite: nil,
		errOnRead:  errors.New("hello, i'm an error :D"),
	}

	repository := NewFileBasedRepository(&mockStorage)
	service := NewService(repository)

	// Act.
	results, err := service.GetAll()

	// Assert.
	assert.EqualError(t, err, expectedErr.Error())
	assert.Nil(t, results)
}

func TestServiceIntegrationStore(t *testing.T) {
	// Arrange.
	expectedDatabase := []domain.Product{
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
			Name:  "Cola Cola 2l",
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
	service := NewService(repository)

	// Act.
	productToCreate := domain.Product{
		ID:    3,
		Name:  "Cola Cola 2l",
		Type:  "Gaseosas",
		Count: 12000,
		Price: 172,
	}

	result, err := service.Store(productToCreate)

	// Assert.
	assert.Nil(t, err)
	assert.Equal(t, expectedDatabase, mockStorage.dataMock)
	assert.Equal(t, productToCreate, result)
}

func TestServiceIntegrationStoreFail(t *testing.T) {
	// Arrange.
	expectedErr := errors.New("hello, i'm a little bug >:(")

	mockStorage := MockStorage{
		dataMock:   nil,
		errOnRead:  nil,
		errOnWrite: errors.New("hello, i'm a little bug >:("),
	}

	repository := NewFileBasedRepository(&mockStorage)
	service := NewService(repository)

	// Act.
	productToCreate := domain.Product{
		ID:    3,
		Name:  "Cola Cola 2l",
		Type:  "Gaseosas",
		Count: 12000,
		Price: 172,
	}

	result, err := service.Store(productToCreate)

	// Assert.
	assert.EqualError(t, err, expectedErr.Error())
	assert.Empty(t, result)
}
