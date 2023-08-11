package repository

import "github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/entity"

type OrderRepositoryInterface interface {
	Save(order *entity.Order) error
	FindById(id string) (*entity.Order, error)
	List() ([]*entity.Order, error)
	// GetTotal() (int, error)
}
