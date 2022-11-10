package product

import (
	"context"
	"errors"

	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Implementacion-Storage/Project/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	Save(ctx context.Context, b domain.Product) (domain.Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	movies, err := s.repo.GetAll(ctx)
	if err != nil {
		return []domain.Product{}, err
	}
	return movies, err
}

func (s *service) Get(ctx context.Context, id int) (product domain.Product, err error) {
	product, err = s.repo.Get(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (s *service) Save(ctx context.Context, p domain.Product) (domain.Product, error) {
	if s.repo.Exists(ctx, p.ID) {
		return domain.Product{}, errors.New("error: movie id already exists")
	}
	id, err := s.repo.Save(ctx, p)
	if err != nil {
		return domain.Product{}, err
	}
	p.ID = int(id)
	return p, nil
}
