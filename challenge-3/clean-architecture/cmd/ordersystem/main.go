package main

import (
	"fmt"

	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/configs"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/event/handler"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/infra/database"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/infra/graph"
	graphserver "github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/infra/graph/server"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/infra/grpc/pb"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/infra/grpc/server"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/infra/grpc/service"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/infra/rabbitmq"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/infra/web/controllers"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/infra/web/webserver"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/pkg/events"
	servicecontainer "github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/service_container"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/setup"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

var (
	conf     *configs.Conf
	db       *database.Database
	rabbitMQ *rabbitmq.RabbitMQ
)

func setupWebserver(input controllers.OrdersControllerInput) *webserver.WebServer {
	webserver := webserver.NewWebServer(conf.WebServerPort)
	ordersController := controllers.NewOrdersController(controllers.OrdersControllerInput{
		CreateOrderUseCase: input.CreateOrderUseCase,
		ListOrdersUseCase:  input.ListOrdersUseCase,
	})
	webserver.AddHandler("/order", "POST", ordersController.CreateOrder)
	webserver.AddHandler("/order", "GET", ordersController.ListOrders)
	return webserver
}

func setupGRPCServer(input service.OrderServiceInput) *server.GRPCServer {
	grpc := server.NewGRPCServer(conf.GRPCServerPort)
	createOrderService := service.NewOrderService(input)
	pb.RegisterOrderServiceServer(grpc.Server, createOrderService)
	reflection.Register(grpc.Server)
	return grpc
}

func setupGraphqlServer(input graphserver.GraphQLServerInput) *graphserver.GraphQLServer {
	graphql := graphserver.NewGraphQLServer(graphserver.GraphQLServerInput{
		Port:      conf.GraphQLServerPort,
		Resolvers: input.Resolvers,
	})
	return graphql
}

func orderCreatedConsumer() {
	msgs, _ := rabbitMQ.Channel.Consume("created_order", "", true, false, false, false, nil)
	for msg := range msgs {
		msg.Ack(false)
		fmt.Printf("Consumer created_order received: %s", string(msg.Body))
	}
}

func init() {
	setup := setup.NewSetup()
	db = setup.DB
	conf = setup.Configs
	rabbitMQ = setup.RabbitMQ
	go orderCreatedConsumer()
}

func main() {
	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQ.Channel,
	})

	createOrderUseCase := servicecontainer.NewCreateOrderUseCase(db.DB, eventDispatcher)
	listOrdersUseCase := servicecontainer.NewListOrdersUseCase(db.DB)

	webserver := setupWebserver(controllers.OrdersControllerInput{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrdersUseCase:  *listOrdersUseCase,
	})
	go webserver.Start()

	grpc := setupGRPCServer(service.OrderServiceInput{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrdersUseCase:  *listOrdersUseCase,
	})
	go grpc.Start()

	graphql := setupGraphqlServer(graphserver.GraphQLServerInput{
		Port: conf.GraphQLServerPort,
		Resolvers: &graph.Resolver{
			CreateOrderUseCase: *createOrderUseCase,
			ListOrdersUseCase:  *listOrdersUseCase,
		},
	})
	go graphql.Start()

	for {
	}
}
