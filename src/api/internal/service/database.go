package service

import (
	"github.com/mnogokotin/golang-packages/database/postgres"
	"github.com/mnogokotin/golang-windows-service/internal/model"
)

func GetEventModelsWithGreaterId(id string) []model.Event {
	postgres, err := postgres.New(postgres.GetConnectionUri())
	if err != nil {
		panic(err)
	}
	defer postgres.Close()

	var eventModels []model.Event
	postgres.Db.Where("id > ?", id).Find(&eventModels)
	return eventModels
}
