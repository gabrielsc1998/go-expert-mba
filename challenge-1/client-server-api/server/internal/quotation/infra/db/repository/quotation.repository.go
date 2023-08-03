package repository

import (
	"context"

	"github.com/gabrielsc1998/go-expert-challenges/client-server-api/server/internal/database"
	quotation "github.com/gabrielsc1998/go-expert-challenges/client-server-api/server/internal/quotation/domain/entity"
	"github.com/gabrielsc1998/go-expert-challenges/client-server-api/server/internal/quotation/infra/db/mapper"
	"github.com/gabrielsc1998/go-expert-challenges/client-server-api/server/internal/quotation/infra/db/models"
)

type QuotationRepository struct {
	db     *database.Database
	mapper *mapper.QuotationMapper
}

func NewQuotationRepository(db *database.Database) *QuotationRepository {
	return &QuotationRepository{db: db, mapper: mapper.NewQuotationMapper()}
}

func (q *QuotationRepository) Save(quotation *quotation.Quotation, ctx *context.Context) error {
	dbWithCtx := q.db.DB.WithContext(*ctx)
	return dbWithCtx.Create(q.mapper.ToModel(quotation)).Error
}

func (q *QuotationRepository) FindById(id string) (*quotation.Quotation, error) {
	quotationModel := &models.QuotationModel{}
	err := q.db.DB.Where("id = ?", id).First(quotationModel).Error
	if err != nil {
		return nil, err
	}
	return q.mapper.ToEntity(quotationModel), nil
}

func (q *QuotationRepository) List() (*[]quotation.Quotation, error) {
	quotationsModels := []models.QuotationModel{}
	err := q.db.DB.Model(&models.QuotationModel{}).Find(&quotationsModels).Error
	if err != nil {
		return nil, err
	}
	quotations := []quotation.Quotation{}
	for _, quotationModel := range quotationsModels {
		quotations = append(quotations, *q.mapper.ToEntity(&quotationModel))
	}
	return &quotations, nil
}
