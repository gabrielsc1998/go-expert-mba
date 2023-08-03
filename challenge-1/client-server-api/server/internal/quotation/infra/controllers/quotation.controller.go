package controllers

import (
	"encoding/json"
	"net/http"

	usecases "github.com/gabrielsc1998/go-expert-challenges/client-server-api/server/internal/quotation/application/use-cases"
)

type QuotationController struct {
	useCase *usecases.GetQuotationUseCase
}

func NewQuotationController(useCase *usecases.GetQuotationUseCase) *QuotationController {
	return &QuotationController{useCase: useCase}
}

func (q *QuotationController) GetQuotation(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	codein := r.URL.Query().Get("codein")

	if code == "" {
		code = "USD"
	}
	if codein == "" {
		codein = "BRL"
	}

	quotation, err := q.useCase.Execute(code, codein)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(quotation)
}
