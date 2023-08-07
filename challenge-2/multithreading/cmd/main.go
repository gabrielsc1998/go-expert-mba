package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gabrielsc1998/go-expert-mba/challenge-2/multithreading/internal/infra/gateway"
)

type GetAddresByCepDto struct {
	Gateway   gateway.CepGatewayInterface
	Cep       string
	Ctx       context.Context
	CtxCancel context.CancelFunc
	Channel   chan map[string]interface{}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	cepFromAPI := make(chan map[string]interface{})

	go getAddressByCep(&GetAddresByCepDto{
		Gateway:   gateway.NewViaCepGateway(),
		Cep:       "83308070",
		Ctx:       ctx,
		CtxCancel: cancel,
		Channel:   cepFromAPI,
	})

	go getAddressByCep(&GetAddresByCepDto{
		Gateway:   gateway.NewBrasilAPIGateway(),
		Cep:       "83308070",
		Ctx:       ctx,
		CtxCancel: cancel,
		Channel:   cepFromAPI,
	})

	go getAddressByCep(&GetAddresByCepDto{
		Gateway:   gateway.NewAPICepGateway(),
		Cep:       "83308070",
		Ctx:       ctx,
		CtxCancel: cancel,
		Channel:   cepFromAPI,
	})

	select {
	case <-ctx.Done():
		if ctx.Err() != nil {
			fmt.Println("Ctx status:", ctx.Err().Error())
		} else {
			fmt.Println("Done!")
		}
	case data := <-cepFromAPI:
		jsonData, _ := json.Marshal(data)
		fmt.Println(string(jsonData))
	}
}

func getAddressByCep(input *GetAddresByCepDto) {
	resp, err := input.Gateway.GetAddressByCep(input.Cep, &input.Ctx)
	if err != nil {
		fmt.Println(input.Gateway.Name(), err)
		return
	}
	m := make(map[string]interface{})
	m["api"] = input.Gateway.Name()
	m["data"] = resp
	input.Channel <- m
	input.CtxCancel()
}
