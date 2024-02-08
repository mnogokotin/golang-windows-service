package database

import (
	"github.com/mnogokotin/golang-packages/database/postgres"
	"github.com/mnogokotin/golang-windows-service/internal/model"
)

func GetEventModelsWithGreaterId(pg *postgres.Postgres, id string) []model.Event {
	var eventModels []model.Event
	pg.Db.Where("id > ?", id).Find(&eventModels)
	return eventModels
}
