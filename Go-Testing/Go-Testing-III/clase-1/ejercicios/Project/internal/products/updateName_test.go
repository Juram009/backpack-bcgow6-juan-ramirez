package products

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockDB struct {
	BeforeUpdate  Product
	ReadWasCalled bool
	AfterUpdate   Product
	db            []Product
}

func (m *MockDB) Write(data interface{}) error {
	Json, _ := json.Marshal(data)
	return json.Unmarshal(Json, &m.db)
}

func (m *MockDB) Read(data interface{}) error {
	m.ReadWasCalled = true
	m.db = []Product{
		{
			Id:           1,
			Name:         "MacBook",
			Color:        "Gris",
			Price:        100000,
			Stock:        7,
			Code:         "8403",
			Published:    false,
			CreationDate: "29/03/2019",
		},
		{
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
	ps, _ := json.Marshal(m.db)
	return json.Unmarshal(ps, &data)
}
func TestUpdateName(t *testing.T) {
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
	newName := "Mackbook Pro"
	db := MockDB{
		BeforeUpdate:  p1,
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
	//Act
	res, _ := repo.UpdateNamePrice(p1.Id, newName, p1.Price)
	//Assert
	assert.NotEqual(t, db.BeforeUpdate, res)
	assert.True(t, db.ReadWasCalled)
	assert.Equal(t, db.AfterUpdate, res)
}
