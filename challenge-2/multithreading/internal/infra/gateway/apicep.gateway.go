package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gabrielsc1998/go-expert-mba/challenge-2/multithreading/internal/domain/entity"
	"github.com/gabrielsc1998/go-expert-mba/challenge-2/multithreading/internal/infra/presenter"
)

type APICepGateway struct {
}

func NewAPICepGateway() *APICepGateway {
	return &APICepGateway{}
}

func (v *APICepGateway) Name() string {
	return "APICep"
}

func (v *APICepGateway) GetAddressByCep(cep string, ctx *context.Context) (*entity.Cep, error) {
	receivedCep := cep[:5] + "-" + cep[5:]
	req, err := http.NewRequestWithContext(*ctx, http.MethodGet, "https://cdn.apicep.com/file/apicep/"+receivedCep+".json", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body := make(map[string]interface{})
		json.NewDecoder(resp.Body).Decode(&body)
		return nil, errors.New(body["message"].(string))
	}

	data := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	apiCepPresenter := presenter.NewAPICepPresenter()
	return apiCepPresenter.Present(data)
}
