package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/usecase"
)

type OrdersController struct {
	createOrderUC usecase.CreateOrderUseCase
	listOrdersUC  usecase.ListOrdersUseCase
}

type OrdersControllerInput struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func NewOrdersController(input OrdersControllerInput) *OrdersController {
	return &OrdersController{
		createOrderUC: input.CreateOrderUseCase,
		listOrdersUC:  input.ListOrdersUseCase,
	}
}

func (c *OrdersController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CreateOrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := c.createOrderUC.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *OrdersController) ListOrders(w http.ResponseWriter, r *http.Request) {
	fmt.Print("ListOrders")
	output, err := c.listOrdersUC.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
