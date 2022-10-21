package products

import "github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Testing/Go-Testing-II/clase-2/clase/internal/domain"

// Service layer for products resource.
type Service interface {
	// GetAll returns all products stored in the system.
	GetAll() ([]domain.Product, error)

	// Store a product in the system.
	Store(productToStore domain.Product) (result domain.Product, err error)
}

type service struct {
	// Repository layer, used for product's data storage.
	Repository Repository
}

// NewService returns a new instance of Service.
func NewService(repository Repository) Service {
	return &service{
		Repository: repository,
	}
}

func (service *service) GetAll() ([]domain.Product, error) {
	return service.Repository.GetAll()
}

func (service *service) Store(productToStore domain.Product) (result domain.Product, err error) {
	return service.Repository.Store(productToStore)
}
