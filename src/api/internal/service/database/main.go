package database

import (
	"github.com/mnogokotin/golang-packages/database/postgres"
	"github.com/mnogokotin/golang-windows-service/internal/model"
)

type Config struct {
	postgresConnectionUri string
}

func NewConfig() *Config {
	return &Config{postgresConnectionUri: "postgres://golang_windows_service_user:golang_windows_service_user_password@host.docker.internal:54350/golang_windows_service?sslmode=disable"}
}

func GetEventModelsWithGreaterId(id string) []model.Event {
	c := NewConfig()
	postgres, err := postgres.New(c.postgresConnectionUri)
	if err != nil {
		panic(err)
	}
	defer postgres.Close()

	var eventModels []model.Event
	postgres.Db.Where("id > ?", id).Find(&eventModels)
	return eventModels
}
