package app

import (
	"github.com/joho/godotenv"
	"github.com/mnogokotin/golang-windows-service/internal/service"
	"os"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}

func Run() {
	outputFilePath := os.Getenv("OUTPUT_FILE_PATH")

	f, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_RDONLY, 0755)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	lastLineId, err := service.GetLastLineId(f)
	if err != nil {
		lastLineId = "0"
	}

	eventModels := service.GetEventModelsWithGreaterId(lastLineId)
	if len(eventModels) > 0 {
		f, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		defer f.Close()
		if err != nil {
			panic(err)
		}

		service.WriteEventsToFile(f, eventModels)
	}
}
