package app

import (
	"encoding/csv"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mnogokotin/golang-packages/database/postgres"
	"github.com/mnogokotin/golang-windows-service/internal/model"
	"io"
	"os"
	"strconv"
	"strings"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}

func Run() {
	separator := '|'

	f, err := os.Open("output.csv")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	lastLine := getLastLine(f)
	lastLineData := strings.Split(lastLine, string(separator))
	lastLineId := lastLineData[0]

	postgres, err := postgres.New(postgres.GetConnectionUri())
	if err != nil {
		panic(err)
	}
	defer postgres.Close()

	var eventModels []model.Event
	postgres.Db.Where("id > ?", lastLineId).Find(&eventModels)

	f, err2 := os.OpenFile("output.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer f.Close()
	if err2 != nil {
		panic(err2)
	}

	w := csv.NewWriter(f)
	w.Comma = separator
	defer w.Flush()

	for _, m := range eventModels {
		if err := w.Write([]string{strconv.Itoa(m.ID), m.Date}); err != nil {
			panic(err)
		}
	}
}

func getLastLine(f *os.File) string {
	line := ""
	var cursor int64 = 0
	stat, _ := f.Stat()
	filesize := stat.Size()
	for {
		cursor -= 1
		f.Seek(cursor, io.SeekEnd)

		char := make([]byte, 1)
		f.Read(char)

		if cursor != -1 && (char[0] == 10 || char[0] == 13) { // stop if we find a line
			break
		}

		line = fmt.Sprintf("%s%s", string(char), line) // there is more efficient way

		if cursor == -filesize { // stop if we are at the begining
			break
		}
	}

	return line
}
