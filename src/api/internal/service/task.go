package service

import (
	"os"
)

func ReadFromDbAndWriteToFile() {
	outputFilePath := os.Getenv("OUTPUT_FILE_PATH")

	f, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_RDONLY, 0755)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	lastLineId, err := GetLastLineId(f)
	if err != nil {
		lastLineId = "0"
	}

	eventModels := GetEventModelsWithGreaterId(lastLineId)
	if len(eventModels) > 0 {
		f, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		defer f.Close()
		if err != nil {
			panic(err)
		}

		WriteEventsToFile(f, eventModels)
	}
}
