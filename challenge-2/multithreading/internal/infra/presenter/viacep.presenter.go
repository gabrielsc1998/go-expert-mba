package presenter

import (
	"encoding/json"

	"github.com/gabrielsc1998/go-expert-mba/challenge-2/multithreading/internal/domain/entity"
)

type ViaCepPresenter struct {
}

type ViaCepResponse struct {
	Code     string `json:"cep"`
	State    string `json:"uf"`
	City     string `json:"localidade"`
	District string `json:"bairro"`
	Address  string `json:"logradouro"`
}

func NewViaCepPresenter() *ViaCepPresenter {
	return &ViaCepPresenter{}
}

func (v *ViaCepPresenter) Present(data map[string]interface{}) (*entity.Cep, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	viaCepResponse := ViaCepResponse{}
	err = json.Unmarshal(jsonData, &viaCepResponse)
	if err != nil {
		return nil, err
	}
	cep := entity.Cep{
		Code:     viaCepResponse.Code,
		State:    viaCepResponse.State,
		City:     viaCepResponse.City,
		District: viaCepResponse.District,
		Address:  viaCepResponse.Address,
	}
	return &cep, nil
}
