package products

import (
	"fmt"

	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Testing/Go-Testing-III/clase-1/ejercicios/Project/pkg/store"
)

type Product struct {
	Id           int    `json:"id"`
	Name         string `json:"name" binding:"required"`
	Color        string `json:"color" binding:"required"`
	Price        int    `json:"price" binding:"required"`
	Stock        int    `json:"stock" binding:"required"`
	Code         string `json:"code" binding:"required"`
	Published    bool   `json:"published" binding:"required"`
	CreationDate string `json:"date" binding:"required"`
}

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

type repository struct {
	db store.Store
} //struct implementa los metodos de la interfaz

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(id int, name, color string, price, stock int, code string, published bool, creationDate string) (Product, error) {
	var products []Product
	err := r.db.Read(&products)
	if err != nil {
		return Product{}, err
	}
	p := Product{id, name, color, price, stock, code, published, creationDate}

	products = append(products, p)
	if err := r.db.Write(products); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) GetAll() ([]Product, error) {
	var products []Product
	err := r.db.Read(&products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) GetOne(id int) (Product, error) {
	var products []Product
	err := r.db.Read(&products)
	if err != nil {
		return Product{}, err
	}
	for i := range products {
		if products[i].Id == id {
			return products[i], nil
		}
	}
	return Product{}, fmt.Errorf("Product %d not found", id)
}

func (r *repository) LastID() (int, error) {
	var products []Product
	err := r.db.Read(&products)
	if err != nil {
		return 0, err
	}
	if len(products) == 0 {
		return 0, nil
	}

	return products[len(products)-1].Id, nil
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
	var products []Product
	err := r.db.Read(&products)
	if err != nil {
		return Product{}, err
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
	if err := r.db.Write(products); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	var products []Product
	var deleted bool
	var pos int
	err := r.db.Read(&products)
	if err != nil {
		return err
	}
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
	if err := r.db.Write(products); err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateNamePrice(id int, name string, price int) (Product, error) {
	var updated bool
	var products []Product
	err := r.db.Read(&products)
	if err != nil {
		return Product{}, err
	}
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
	if err := r.db.Write(products); err != nil {
		return Product{}, err
	}
	return product, nil
}
