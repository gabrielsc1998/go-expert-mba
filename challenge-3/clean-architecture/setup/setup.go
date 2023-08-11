package setup

import (
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/configs"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/infra/database"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/infra/rabbitmq"
)

type Setup struct {
	DB       *database.Database
	Configs  *configs.Conf
	RabbitMQ *rabbitmq.RabbitMQ
}

func NewSetup() *Setup {
	configs := loadConfigs()
	db := connectDatabase(configs)
	rabbitmq := connectRabbitMQ(configs)

	return &Setup{
		Configs:  configs,
		DB:       db,
		RabbitMQ: rabbitmq,
	}
}

func loadConfigs() *configs.Conf {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	return configs
}

func connectDatabase(configs *configs.Conf) *database.Database {
	db := database.NewDatabase()
	dbOptions := database.DatabaseOptions{
		Driver:   configs.DBDriver,
		User:     configs.DBUser,
		Password: configs.DBPassword,
		Host:     configs.DBHost,
		Port:     configs.DBPort,
		Name:     configs.DBName,
	}
	db.Connect(dbOptions)
	return db
}

func connectRabbitMQ(configs *configs.Conf) *rabbitmq.RabbitMQ {
	rabbitMQ := rabbitmq.NewRabbitMQ()
	err := rabbitMQ.Connect(rabbitmq.RabbitMQOptions{
		User:     configs.RabbitMQUser,
		Password: configs.RabbitMQPassword,
		Host:     configs.RabbitMQHost,
		Port:     configs.RabbitMQPort,
	})
	if err != nil {
		panic(err)
	}
	return rabbitMQ
}
