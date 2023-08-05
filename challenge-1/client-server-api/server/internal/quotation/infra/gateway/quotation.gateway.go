package gateway

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gabrielsc1998/go-expert-mba/challenge-1/client-server-api/server/configs"
	quotation "github.com/gabrielsc1998/go-expert-mba/challenge-1/client-server-api/server/internal/quotation/domain/entity"
)

type EconomyAwesomeAPIGateway struct {
}

func NewEconomyAwesomeAPIGateway() *EconomyAwesomeAPIGateway {
	return &EconomyAwesomeAPIGateway{}
}

func (g *EconomyAwesomeAPIGateway) GetQuotation(code string, codeIn string) (*quotation.Quotation, error) {
	config, err := configs.LoadConfig()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	url := config.QuotationAPIUrl + "/" + code + "-" + codeIn
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)

	if ctx.Err() == context.DeadlineExceeded {
		log.Println("Timeout exceeded for request to " + url)
	}

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	expectedCode := code + codeIn
	quotation, err := g.parseQuotation(resp, expectedCode)
	if err != nil {
		return nil, err
	}
	return quotation, nil
}

func (g *EconomyAwesomeAPIGateway) parseQuotation(resp *http.Response, code string) (*quotation.Quotation, error) {
	bodyMap := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&bodyMap)
	quotation := quotation.Quotation{}
	jsonData, _ := json.Marshal(bodyMap[code])
	json.Unmarshal(jsonData, &quotation)
	return &quotation, nil
}
