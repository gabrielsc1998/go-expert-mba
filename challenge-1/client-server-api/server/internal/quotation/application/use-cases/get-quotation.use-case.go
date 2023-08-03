package usecases

import (
	"context"
	"log"
	"time"

	"github.com/gabrielsc1998/go-expert-challenges/client-server-api/server/internal/quotation/infra/db/repository"
	"github.com/gabrielsc1998/go-expert-challenges/client-server-api/server/internal/quotation/infra/gateway"
	"github.com/gabrielsc1998/go-expert-challenges/client-server-api/server/internal/quotation/infra/presenter"
)

type GetQuotationUseCase struct {
	repository *repository.QuotationRepository
	gtw        *gateway.EconomyAwesomeAPIGateway
}

func NewGetQuotationUseCase(repository *repository.QuotationRepository, gtw *gateway.EconomyAwesomeAPIGateway) *GetQuotationUseCase {
	return &GetQuotationUseCase{repository: repository, gtw: gtw}
}

func (uc *GetQuotationUseCase) Execute(code string, codeIn string) (*presenter.QuotationPresenter, error) {
	receivedQuotation, err := uc.gtw.GetQuotation(code, codeIn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	err = uc.repository.Save(receivedQuotation, &ctx)
	defer cancel()

	if ctx.Err() == context.DeadlineExceeded {
		log.Println("Timeout exceeded for save quotation")
	}

	if err != nil {
		return nil, err
	}
	return receivedQuotation.ToPresent(), nil
}
