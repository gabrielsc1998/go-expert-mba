package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gabrielsc1998/go-expert-mba/challenge-2/multithreading/internal/infra/gateway"
)

func main() {
	viaCepGateway := gateway.NewViaCepGateway()
	brasilAPIGateway := gateway.NewViaCepGateway()
	apiCepGateway := gateway.NewAPICepGateway()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	cepFromAPI := make(chan map[string]interface{})

	go func() {
		resp, err := viaCepGateway.GetAddressByCep("83308070", &ctx)
		if err != nil {
			fmt.Println("ViaCep:", err)
			return
		}
		m := make(map[string]interface{})
		m["api"] = "ViaCep"
		m["data"] = resp
		cepFromAPI <- m
		cancel()
	}()

	go func() {
		resp, err := apiCepGateway.GetAddressByCep("83308070", &ctx)
		if err != nil {
			fmt.Println("APICep:", err)
			return
		}
		m := make(map[string]interface{})
		m["api"] = "APICep"
		m["data"] = resp
		cepFromAPI <- m
		cancel()
	}()

	go func() {
		resp, err := brasilAPIGateway.GetAddressByCep("83308070", &ctx)
		if err != nil {
			fmt.Println("BrasilAPI:", err)
			return
		}
		m := make(map[string]interface{})
		m["api"] = "BrasilAPI"
		m["data"] = resp
		cepFromAPI <- m
		cancel()
	}()

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
