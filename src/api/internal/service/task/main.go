package task

import (
	"github.com/mnogokotin/golang-windows-service/internal/service/csv"
	"github.com/mnogokotin/golang-windows-service/internal/service/database"
	"os"
)

type Config struct {
	outputFilePath string
}

func newConfig() *Config {
	return &Config{outputFilePath: "C:\\vhosts\\output.csv"}
}

func ReadFromDbAndWriteToFile() {
	c := newConfig()

	f, err := os.OpenFile(c.outputFilePath, os.O_CREATE|os.O_RDONLY, 0755)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	lastLineId, err := csv.GetLastLineId(f)
	if err != nil {
		lastLineId = "0"
	}

	eventModels := database.GetEventModelsWithGreaterId(lastLineId)
	if len(eventModels) > 0 {
		f, err := os.OpenFile(c.outputFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		defer f.Close()
		if err != nil {
			panic(err)
		}

		csv.WriteEventsToFile(f, eventModels)
	}
}
