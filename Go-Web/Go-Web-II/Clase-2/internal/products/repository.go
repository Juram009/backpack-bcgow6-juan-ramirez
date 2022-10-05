package products

import (
	"fmt"
)

type Product struct {
	Id           int    `json:"id" binding:"required"`
	Name         string `json:"nombre" binding:"required"`
	Color        string `json:"color" binding:"required"`
	Price        int    `json:"precio" binding:"required"`
	Stock        int    `json:"stock" binding:"required"`
	Code         string `json:"codigo" binding:"required"`
	Published    bool   `json:"publicado" binding:"required"`
	CreationDate string `json:"fecha" binding:"required"`
}

var ps []Product
var lastID int

// ***Importado**//
type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name, color string, price, stock int, code string, published bool, creationDate string) (Product, error)
	LastID() (int, error)
	Update(id int, name, color string, price, stock int, code string, published bool, creationDate string) (Product, error)
}

type repository struct{} //struct implementa los metodos de la interfaz

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Store(id int, name, color string, price, stock int, code string, published bool, creationDate string) (Product, error) {
	p := Product{id, name, color, price, stock, code, published, creationDate}
	ps = append(ps, p)
	lastID = p.Id
	return p, nil
}

func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Update(id int, name, color string, price, stock int, code string, published bool, creationDate string) (Product, error) {
	p := Product{Name: name,
		Color:        color,
		Price:        price,
		Stock:        stock,
		Code:         code,
		Published:    published,
		CreationDate: creationDate,
	}
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Product %d not found", id)
	}
	return p, nil
}
