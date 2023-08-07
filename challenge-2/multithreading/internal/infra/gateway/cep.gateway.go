package gateway

import (
	"context"

	"github.com/gabrielsc1998/go-expert-mba/challenge-2/multithreading/internal/domain/entity"
)

type CepGatewayInterface interface {
	Name() string
	GetAddressByCep(cep string, ctx *context.Context) (*entity.Cep, error)
}
