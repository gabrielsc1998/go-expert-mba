package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gabrielsc1998/go-expert-mba/challenge-2/multithreading/internal/domain/entity"
	"github.com/gabrielsc1998/go-expert-mba/challenge-2/multithreading/internal/infra/presenter"
)

type ViaCepGateway struct {
}

func NewViaCepGateway() *ViaCepGateway {
	return &ViaCepGateway{}
}

func (v *ViaCepGateway) Name() string {
	return "ViaCep"
}

func (v *ViaCepGateway) GetAddressByCep(cep string, ctx *context.Context) (*entity.Cep, error) {
	req, err := http.NewRequestWithContext(*ctx, http.MethodGet, "https://viacep.com.br/ws/"+cep+"/json/", nil)
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
	viaCepPresenter := presenter.NewViaCepPresenter()
	return viaCepPresenter.Present(data)
}
