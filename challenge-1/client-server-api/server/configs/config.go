package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type conf struct {
	ServerPort      string
	QuotationAPIUrl string
}

func LoadConfig() (*conf, error) {
	err := godotenv.Load("../.env")
	return &conf{
		ServerPort:      os.Getenv("SERVER_PORT"),
		QuotationAPIUrl: os.Getenv("QUOTATION_API_URL"),
	}, err
}
