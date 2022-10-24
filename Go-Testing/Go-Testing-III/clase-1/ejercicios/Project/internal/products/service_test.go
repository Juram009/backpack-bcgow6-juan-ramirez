package products

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestProducts() []Product {
	return []Product{{
		Id:           1,
		Name:         "MacBook",
		Color:        "Gris",
		Price:        100000,
		Stock:        7,
		Code:         "8403",
		Published:    false,
		CreationDate: "29/03/2019",
	}, {
		Id:           2,
		Name:         "Pantalla",
		Color:        "Negra",
		Price:        1000,
		Stock:        2,
		Code:         "2639",
		Published:    true,
		CreationDate: "09/03/2020",
	},
	}
}
func TestUpdate(t *testing.T) {
	//Arrange
	products := getTestProducts()
	expectedProducts := []Product{products[0], {
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
	_, errPut := service.Update(p3.Id, p3.Name, p3.Color, p3.Price, p3.Stock, p3.Code, p3.Published, p3.CreationDate)
	res, errGet := service.GetAll()

	//Assert
	assert.Equal(t, expectedProducts, res)
	assert.Nil(t, errPut)
	assert.Nil(t, errGet)
	assert.True(t, mockStorage.ReadWasCalled)
}
func TestFailUpdateOnRead(t *testing.T) {
	//Arrange
	products := getTestProducts()
	mockStorage := MockStorage{
		dataMock:  products,
		errOnRead: fmt.Errorf("Error on read"),
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
	_, err := service.Update(p3.Id, p3.Name, p3.Color, p3.Price, p3.Stock, p3.Code, p3.Published, p3.CreationDate)

	//Assert
	assert.NotNil(t, err)
}
func TestFailUpdateOnWrite(t *testing.T) {
	//Arrange
	products := getTestProducts()
	mockStorage := MockStorage{
		dataMock:   products,
		errOnWrite: fmt.Errorf("Error on write"),
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
	_, err := service.Update(p3.Id, p3.Name, p3.Color, p3.Price, p3.Stock, p3.Code, p3.Published, p3.CreationDate)

	//Assert
	assert.NotNil(t, err)
}
func TestFailUpdateOnUpdate(t *testing.T) {
	products := getTestProducts()
	newName := "Mackbook"
	db := MockStorage{
		dataMock: products,
	}
	repo := NewRepository(&db)
	service := NewService(repo)
	p3 := Product{
		Id:           17,
		Name:         "Pantall",
		Color:        "Negra",
		Price:        1000,
		Stock:        2,
		Code:         "2639",
		Published:    true,
		CreationDate: "09/03/2020",
	}
	//Act
	_, err := service.Update(p3.Id, newName, p3.Color, p3.Price, p3.Stock, p3.Code, p3.Published, p3.CreationDate)
	//Assert
	assert.NotNil(t, err)
}

func TestDelete(t *testing.T) {
	//Arrange
	products := getTestProducts()
	expectedProducts := []Product{products[0]}
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
func TestFailDeleteOnRead(t *testing.T) {
	//Arrange
	products := getTestProducts()

	mockStorage := MockStorage{
		dataMock:  products,
		errOnRead: fmt.Errorf("Error on read"),
	}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	//Act
	err := service.Delete(1)

	//Assert
	assert.NotNil(t, err)
}
func TestFailDeleteOnWrite(t *testing.T) {
	//Arrange
	products := getTestProducts()

	mockStorage := MockStorage{
		dataMock:   products,
		errOnWrite: fmt.Errorf("Error on Write"),
	}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	//Act
	err := service.Delete(1)

	//Assert
	assert.NotNil(t, err)
}

func TestFailOnRead(t *testing.T) {
	//Arrange
	products := getTestProducts()
	expectedErr := fmt.Errorf("Error reading the database")
	mockStorage := MockStorage{
		dataMock:  products,
		errOnRead: fmt.Errorf("Error reading the database"),
	}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	//Act
	_, err := service.GetAll()

	//Assert
	assert.EqualError(t, err, expectedErr.Error())
}

func TestFailOnWrite(t *testing.T) {
	//Arrange
	products := getTestProducts()
	p1 := products[0]
	expectedErr := fmt.Errorf("Error writting the database")
	mockStorage := MockStorage{
		dataMock:   products,
		errOnWrite: fmt.Errorf("Error writting the database"),
	}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	//Act
	_, err := service.Store(p1.Name, p1.Color, p1.Price, p1.Stock, p1.Code, p1.Published, p1.CreationDate)

	//Assert
	assert.EqualError(t, err, expectedErr.Error())
}

func TestGetOne(t *testing.T) {
	//Arrange
	products := getTestProducts()
	expectedProduct := products[0]
	//expectedErr := fmt.Errorf("Error getting the product")
	mockStorage := MockStorage{
		dataMock: products,
		//errOnWrite: fmt.Errorf("Error writting the database"),
	}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	//Act
	res, err := service.GetOne(products[0].Id)

	//Assert
	assert.Equal(t, expectedProduct, res)
	assert.Nil(t, err)
	assert.True(t, mockStorage.ReadWasCalled)
}
func TestFailGetOne(t *testing.T) {
	//Arrange
	products := getTestProducts()
	mockStorage := MockStorage{
		dataMock: products,
	}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	//Act
	res, err := service.GetOne(7)

	//Assert
	assert.Equal(t, Product{}, res)
	assert.NotNil(t, err)
}
func TestFailGetOneOnRead(t *testing.T) {
	//Arrange
	products := getTestProducts()
	mockStorage := MockStorage{
		dataMock:  products,
		errOnRead: fmt.Errorf("Error on reading"),
	}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	//Act
	res, err := service.GetOne(7)

	//Assert
	assert.Equal(t, Product{}, res)
	assert.NotNil(t, err)
}
func TestUpdateNamePrice(t *testing.T) {
	products := getTestProducts()
	newName := "Mackbook Pro"
	db := MockDB{
		BeforeUpdate:  products[0],
		ReadWasCalled: false,
		db:            products,
		AfterUpdate: Product{
			Id:           1,
			Name:         newName,
			Color:        "Gris",
			Price:        100000,
			Stock:        7,
			Code:         "8403",
			Published:    false,
			CreationDate: "29/03/2019",
		},
	}
	repo := NewRepository(&db)
	service := NewService(repo)
	//Act
	res, _ := service.UpdateNamePrice(products[0].Id, newName, products[0].Price)
	//Assert
	assert.NotEqual(t, db.BeforeUpdate, res)
	assert.True(t, db.ReadWasCalled)
	assert.Equal(t, db.AfterUpdate, res)
}
func TestFailUpdateNamePriceOnRead(t *testing.T) {
	products := getTestProducts()
	newName := "Mackbook"
	db := MockStorage{
		dataMock:  products,
		errOnRead: fmt.Errorf("Error on read"),
	}
	repo := NewRepository(&db)
	service := NewService(repo)
	//Act
	_, err := service.UpdateNamePrice(products[0].Id, newName, products[0].Price)
	//Assert
	assert.NotNil(t, err)
}
func TestFailUpdateNamePriceOnUpdate(t *testing.T) {
	products := getTestProducts()
	newName := "Mackbook"
	db := MockStorage{
		dataMock: products,
	}
	repo := NewRepository(&db)
	service := NewService(repo)
	//Act
	_, err := service.UpdateNamePrice(17, newName, products[0].Price)
	//Assert
	assert.NotNil(t, err)
}
func TestFailUpdateNamePriceOnWrite(t *testing.T) {
	products := getTestProducts()
	newName := "Mackbook"
	db := MockStorage{
		dataMock:   products,
		errOnWrite: fmt.Errorf("Error on write"),
	}
	repo := NewRepository(&db)
	service := NewService(repo)
	//Act
	_, err := service.UpdateNamePrice(products[0].Id, newName, products[0].Price)
	//Assert
	assert.NotNil(t, err)
}
func TestFailStoreOnRead(t *testing.T) {
	//Arrange
	products := getTestProducts()
	p1 := products[0]
	expectedErr := fmt.Errorf("Error reading the database")
	mockStorage := MockStorage{
		dataMock:  products,
		errOnRead: fmt.Errorf("Error reading the database"),
	}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	//Act
	res, err := service.Store(p1.Name, p1.Color, p1.Price, p1.Stock, p1.Code, p1.Published, p1.CreationDate)

	//Assert
	assert.Equal(t, Product{}, res)
	assert.EqualError(t, err, expectedErr.Error())
}
func TestStore(t *testing.T) {
	//Arrange
	products := getTestProducts()
	p3 := Product{
		Id:           3,
		Name:         "name",
		Color:        "color",
		Price:        12,
		Stock:        45,
		Code:         "code",
		Published:    true,
		CreationDate: "creationDate",
	}
	mockStorage := MockStorage{
		dataMock: products,
	}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	//Act
	res, err := service.Store("name", "color", 12, 45, "code", true, "creationDate")

	//Assert
	assert.Equal(t, p3, res)
	assert.Nil(t, err)
}
func TestLastIdError(t *testing.T) {
	//Arrange
	products := []Product{}
	mockStorage := MockStorage{
		dataMock: products,
	}
	repo := NewRepository(&mockStorage)
	//Act
	res, _ := repo.LastID()
	//Assert
	assert.Equal(t, 0, res)
}
