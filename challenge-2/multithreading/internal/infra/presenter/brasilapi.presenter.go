package presenter

import (
	"encoding/json"

	"github.com/gabrielsc1998/go-expert-mba/challenge-2/multithreading/internal/domain/entity"
)

type BrasilAPIPresenter struct {
}

type BrasilAPIResponse struct {
	Code     string `json:"cep"`
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"neighborhood"`
	Address  string `json:"street"`
}

func NewBrasilAPIPresenter() *BrasilAPIPresenter {
	return &BrasilAPIPresenter{}
}

func (v *BrasilAPIPresenter) Present(data map[string]interface{}) (*entity.Cep, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	BrasilAPIResponse := BrasilAPIResponse{}
	err = json.Unmarshal(jsonData, &BrasilAPIResponse)
	if err != nil {
		return nil, err
	}
	cep := entity.Cep{
		Code:     BrasilAPIResponse.Code,
		State:    BrasilAPIResponse.State,
		City:     BrasilAPIResponse.City,
		District: BrasilAPIResponse.District,
		Address:  BrasilAPIResponse.Address,
	}
	return &cep, nil
}
