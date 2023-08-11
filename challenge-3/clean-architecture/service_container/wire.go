//go:build wireinject
// +build wireinject

package servicecontainer

import (
	"database/sql"

	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/event"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/infra/database/repository"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/usecase"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	repository.NewOrderRepository,
	wire.Bind(new(repository.OrderRepositoryInterface), new(*repository.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewListOrdersUseCase(db *sql.DB) *usecase.ListOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewListOrdersUseCase,
	)
	return &usecase.ListOrdersUseCase{}
}
