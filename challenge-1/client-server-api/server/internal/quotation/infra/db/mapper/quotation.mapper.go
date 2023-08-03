package mapper

import (
	quotation "github.com/gabrielsc1998/go-expert-challenges/client-server-api/server/internal/quotation/domain/entity"
	"github.com/gabrielsc1998/go-expert-challenges/client-server-api/server/internal/quotation/infra/db/models"
)

type QuotationMapper struct {
}

func NewQuotationMapper() *QuotationMapper {
	return &QuotationMapper{}
}

func (q *QuotationMapper) ToModel(quotation *quotation.Quotation) *models.QuotationModel {
	return &models.QuotationModel{
		Id:         quotation.Id,
		Code:       quotation.Code,
		Codein:     quotation.Codein,
		Name:       quotation.Name,
		High:       quotation.High,
		Low:        quotation.Low,
		VarBid:     quotation.VarBid,
		PctChange:  quotation.PctChange,
		Bid:        quotation.Bid,
		Ask:        quotation.Ask,
		Timestamp:  quotation.Timestamp,
		CreateDate: quotation.CreateDate,
	}
}

func (q *QuotationMapper) ToEntity(model *models.QuotationModel) *quotation.Quotation {
	return &quotation.Quotation{
		Id:         model.Id,
		Code:       model.Code,
		Codein:     model.Codein,
		Name:       model.Name,
		High:       model.High,
		Low:        model.Low,
		VarBid:     model.VarBid,
		PctChange:  model.PctChange,
		Bid:        model.Bid,
		Ask:        model.Ask,
		Timestamp:  model.Timestamp,
		CreateDate: model.CreateDate,
	}
}
