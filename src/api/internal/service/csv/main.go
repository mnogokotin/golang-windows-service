package csv

import (
	"encoding/csv"
	ufile "github.com/mnogokotin/golang-packages/utils/file"
	"github.com/mnogokotin/golang-windows-service/internal/model"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	csvSeparator string
}

func newConfig() *Config {
	return &Config{csvSeparator: "|"}
}

func GetLastLineId(file *os.File) (string, error) {
	lastLine, err := ufile.GetLastLine(file)
	if err != nil {
		return "", err
	}
	return getIdFromLine(lastLine), nil
}

func getIdFromLine(line string) string {
	c := newConfig()
	lineData := strings.Split(line, c.csvSeparator)
	return lineData[0]
}

func WriteEventsToFile(file *os.File, eventModels []model.Event) {
	c := newConfig()
	w := csv.NewWriter(file)
	w.Comma = []rune(c.csvSeparator)[0]
	defer w.Flush()

	for _, m := range eventModels {
		if err := w.Write([]string{strconv.Itoa(m.ID), m.Date}); err != nil {
			panic(err)
		}
	}
}
