package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Quotation struct {
	Value string `json:"value"`
}

const BASE_DIR = "../data"
const FILE_PATH = BASE_DIR + "/cotacao.txt"

func init() {
	dirExists, _ := os.Stat(BASE_DIR)
	if dirExists == nil {
		err := os.Mkdir(BASE_DIR, 0777)
		if err != nil {
			panic(err)
		}
	}
	fileExists, _ := os.Stat(FILE_PATH)
	if fileExists == nil {
		f, err := os.Create(FILE_PATH)
		if err != nil {
			panic(err)
		}
		defer f.Close()
	}
}

func main() {
	quotation := getQuotation()
	if quotation.Value != "" {
		saveQuotation(quotation)
	}
}

func getQuotation() *Quotation {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	url := "http://localhost:8080/cotacao"
	// url := "http://localhost:8080/cotacao?code=USD&codein=ARS"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("Timeout exceeded for request to " + url)
	}
	if err != nil {
		panic(err)
	}
	return parseQuotation(resp)
}

func parseQuotation(respone *http.Response) *Quotation {
	defer respone.Body.Close()
	bodyMap := make(map[string]interface{})
	json.NewDecoder(respone.Body).Decode(&bodyMap)
	quotation := Quotation{}
	jsonData, _ := json.Marshal(bodyMap)
	json.Unmarshal(jsonData, &quotation)
	return &quotation
}

func saveQuotation(quotation *Quotation) {
	f, err := os.OpenFile(FILE_PATH, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString("Dólar: " + quotation.Value + "\n")
	if err != nil {
		panic(err)
	}
}
