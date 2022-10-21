package products

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubDB struct {
	db []Product
}

func (s *StubDB) Write(data interface{}) error {
	return nil
}

func (s *StubDB) Read(data interface{}) error {
	s.db = []Product{
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
	ps, _ := json.Marshal(s.db)
	return json.Unmarshal(ps, &data)
}

func TestGetAll(t *testing.T) {
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
	db := StubDB{products}
	repo := NewRepository(&db)
	//Act
	res, _ := repo.GetAll()
	//Assert
	assert.Equal(t, products, res)
}
