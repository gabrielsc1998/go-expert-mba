package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gabrielsc1998/go-expert-mba/challenge-2/multithreading/internal/domain/entity"
	"github.com/gabrielsc1998/go-expert-mba/challenge-2/multithreading/internal/infra/presenter"
)

type BrasilAPIGateway struct {
}

func NewBrasilAPIGateway() *BrasilAPIGateway {
	return &BrasilAPIGateway{}
}

func (v *BrasilAPIGateway) Name() string {
	return "BrasilAPI"
}

func (v *BrasilAPIGateway) GetAddressByCep(cep string, ctx *context.Context) (*entity.Cep, error) {
	req, err := http.NewRequestWithContext(*ctx, http.MethodGet, "https://brasilapi.com.br/api/cep/v1/"+cep, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error")
	}

	data := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	BrasilAPIPresenter := presenter.NewBrasilAPIPresenter()
	return BrasilAPIPresenter.Present(data)
}
