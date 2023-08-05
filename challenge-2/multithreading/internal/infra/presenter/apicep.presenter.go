package presenter

import (
	"encoding/json"

	"github.com/gabrielsc1998/go-expert-mba/challenge-2/multithreading/internal/domain/entity"
)

type APICepPresenter struct {
}

type APICepResponse struct {
	Code     string `json:"code"`
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"district"`
	Address  string `json:"address"`
}

func NewAPICepPresenter() *APICepPresenter {
	return &APICepPresenter{}
}

func (v *APICepPresenter) Present(data map[string]interface{}) (*entity.Cep, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	apiCepResponse := APICepResponse{}
	err = json.Unmarshal(jsonData, &apiCepResponse)
	if err != nil {
		return nil, err
	}
	cep := entity.Cep{
		Code:     apiCepResponse.Code,
		State:    apiCepResponse.State,
		City:     apiCepResponse.City,
		District: apiCepResponse.District,
		Address:  apiCepResponse.Address,
	}
	return &cep, nil
}
