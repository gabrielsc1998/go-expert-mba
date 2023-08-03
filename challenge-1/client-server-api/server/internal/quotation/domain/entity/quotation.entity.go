package quotation

import (
	"github.com/gabrielsc1998/go-expert-challenges/client-server-api/server/internal/quotation/infra/presenter"
)

type Quotation struct {
	Id         string `json:"id"`
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func (q *Quotation) ToPresent() *presenter.QuotationPresenter {
	return &presenter.QuotationPresenter{
		Value: q.Bid,
	}
}
