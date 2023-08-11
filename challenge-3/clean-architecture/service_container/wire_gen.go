// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package servicecontainer

import (
	"database/sql"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/event"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/infra/database/repository"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/usecase"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/pkg/events"
	"github.com/google/wire"
)

// Injectors from wire.go:

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	orderRepository := repository.NewOrderRepository(db)
	orderCreated := event.NewOrderCreated()
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository, orderCreated, eventDispatcher)
	return createOrderUseCase
}

func NewListOrdersUseCase(db *sql.DB) *usecase.ListOrdersUseCase {
	orderRepository := repository.NewOrderRepository(db)
	listOrdersUseCase := usecase.NewListOrdersUseCase(orderRepository)
	return listOrdersUseCase
}

// wire.go:

var setOrderRepositoryDependency = wire.NewSet(repository.NewOrderRepository, wire.Bind(new(repository.OrderRepositoryInterface), new(*repository.OrderRepository)))

var setEventDispatcherDependency = wire.NewSet(events.NewEventDispatcher, event.NewOrderCreated, wire.Bind(new(events.EventInterface), new(*event.OrderCreated)), wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)))

var setOrderCreatedEvent = wire.NewSet(event.NewOrderCreated, wire.Bind(new(events.EventInterface), new(*event.OrderCreated)))