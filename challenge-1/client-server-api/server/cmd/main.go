package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/gabrielsc1998/go-expert-mba/challenge-1/client-server-api/server/configs"
	"github.com/gabrielsc1998/go-expert-mba/challenge-1/client-server-api/server/internal/database"
	usecases "github.com/gabrielsc1998/go-expert-mba/challenge-1/client-server-api/server/internal/quotation/application/use-cases"
	"github.com/gabrielsc1998/go-expert-mba/challenge-1/client-server-api/server/internal/quotation/infra/controllers"
	"github.com/gabrielsc1998/go-expert-mba/challenge-1/client-server-api/server/internal/quotation/infra/db/models"
	"github.com/gabrielsc1998/go-expert-mba/challenge-1/client-server-api/server/internal/quotation/infra/db/repository"
	"github.com/gabrielsc1998/go-expert-mba/challenge-1/client-server-api/server/internal/quotation/infra/gateway"
)

func initializeDatabase() *database.Database {
	db := database.NewDatabase()
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	db.DB.AutoMigrate(&models.QuotationModel{})
	return db
}

func main() {
	configs, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	db := initializeDatabase()
	gateway := gateway.NewEconomyAwesomeAPIGateway()
	repository := repository.NewQuotationRepository(db)
	// x, err := repository.List()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Print(x)
	getQuotationUseCase := usecases.NewGetQuotationUseCase(repository, gateway)
	quotationHandler := controllers.NewQuotationController(getQuotationUseCase)
	r := chi.NewRouter()
	r.Route("/cotacao", func(r chi.Router) {
		r.Get("/", quotationHandler.GetQuotation)
	})

	http.ListenAndServe(":"+configs.ServerPort, r)
}
