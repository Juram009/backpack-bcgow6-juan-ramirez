package products

import (
	"fmt"
)

type Product struct {
	Id           int    `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Color        string `json:"color" binding:"required"`
	Price        int    `json:"price" binding:"required"`
	Stock        int    `json:"stock" binding:"required"`
	Code         string `json:"code" binding:"required"`
	Published    bool   `json:"published" binding:"required"`
	CreationDate string `json:"date" binding:"required"`
}

var products []Product
var lastID int

// ***Importado**//
type Repository interface {
	GetAll() ([]Product, error)
	GetOne(id int) (Product, error)
	Store(id int, name, color string, price, stock int, code string, published bool, creationDate string) (Product, error)
	LastID() (int, error)
	Update(id int, name, color string, price, stock int, code string, published bool, creationDate string) (Product, error)
	Delete(id int) error
	UpdateNamePrice(id int, name string, price int) (Product, error)
}

type repository struct{} //struct implementa los metodos de la interfaz

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Store(id int, name, color string, price, stock int, code string, published bool, creationDate string) (Product, error) {
	p := Product{id, name, color, price, stock, code, published, creationDate}
	products = append(products, p)
	lastID = p.Id
	return p, nil
}

func (r *repository) GetAll() ([]Product, error) {
	return products, nil
}

func (r *repository) GetOne(id int) (Product, error) {
	for i := range products {
		if products[i].Id == id {
			return products[i], nil
		}
	}
	return Product{}, fmt.Errorf("Product %d not found", id)
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
	for i := range products {
		if products[i].Id == id {
			p.Id = id
			products[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Product %d not found", id)
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	var deleted bool
	var pos int
	for i := range products {
		if products[i].Id == id {
			pos = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("product id %d not exists", id)
	}

	products = append(products[:pos], products[pos+1:]...)
	return nil
}

func (r *repository) UpdateNamePrice(id int, name string, price int) (Product, error) {
	var updated bool
	product := Product{}
	for i := range products {
		if products[i].Id == id {
			products[i].Name = name
			products[i].Price = price
			product = products[i]
			updated = true
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("Product with id %d not exists", id)
	}
	return product, nil
}
