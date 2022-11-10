package products

import (
	"context"
	"fmt"

	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/storage/internal/domain"
)

type Service interface {
	Get(context.Context, string) (domain.Product, error)
	GetAll(context.Context) ([]domain.Product, error)
	Store(context.Context, domain.Product) (int, error)
	Update(context.Context, domain.Product) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) Get(ctx context.Context, name string) (domain.Product, error) {
	return s.repository.GetByName(ctx, name)
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Store(ctx context.Context, p domain.Product) (int, error) {
	return s.repository.Store(ctx, p)
}

func (s *service) Update(ctx context.Context, p domain.Product) error {
	if !s.repository.Exists(ctx, p.ID) {
		return fmt.Errorf("not exists product id %v", p.ID)
	}

	product, err := s.repository.GetByName(ctx, p.Name)
	if err != nil {
		return err
	}

	if p.Name == "" {
		p.Name = product.Name
	}

	if p.Count == 0 {
		p.Count = product.Count
	}

	if p.Type == "" {
		p.Type = product.Type
	}

	if p.Price == 0 {
		p.Price = product.Price
	}

	return s.repository.Update(ctx, p)
}
