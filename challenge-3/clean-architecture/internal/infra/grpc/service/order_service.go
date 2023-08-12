package service

import (
	"context"

	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/infra/grpc/pb"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

type OrderServiceInput struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func NewOrderService(input OrderServiceInput) *OrderService {
	return &OrderService{
		CreateOrderUseCase: input.CreateOrderUseCase,
		ListOrdersUseCase:  input.ListOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.CreateOrderInputDTO{
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.XBlank) (*pb.ListOrdersResponse, error) {
	output, err := s.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	orders := []*pb.Order{}
	for _, order := range output {
		orders = append(orders, &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}
	return &pb.ListOrdersResponse{
		Orders: orders,
	}, nil
}

func (s *OrderService) ListOrdersStream(in *pb.XBlank, stream pb.OrderService_ListOrdersStreamServer) error {
	output, err := s.ListOrdersUseCase.Execute()
	if err != nil {
		return err
	}
	for _, order := range output {
		outOrder := &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}
		err := stream.Send(outOrder)
		if err != nil {
			return err
		}
	}
	return nil
}
