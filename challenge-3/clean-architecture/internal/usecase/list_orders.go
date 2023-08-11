package usecase

import (
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/infra/database/repository"
)

type ListOrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository repository.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository repository.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() ([]ListOrdersOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	dto := []ListOrdersOutputDTO{}
	if err != nil {
		return dto, err
	}
	for order := range orders {
		orders[order].CalculateFinalPrice()
		dto = append(dto, ListOrdersOutputDTO{
			ID:         orders[order].ID,
			Price:      orders[order].Price,
			Tax:        orders[order].Tax,
			FinalPrice: orders[order].FinalPrice,
		})
	}
	return dto, nil
}
