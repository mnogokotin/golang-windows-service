package csv

import (
	"fmt"
	ufile "github.com/mnogokotin/golang-packages/utils/file"
	"os"
	"reflect"
	"strings"
)

func GetLastLineId(file *os.File, csvSeparator string) (string, error) {
	lastLine, err := ufile.GetLastLine(file)
	if err != nil {
		return "", err
	}
	return getIdFromLine(lastLine, csvSeparator), nil
}

func getIdFromLine(line string, csvSeparator string) string {
	lineData := strings.Split(line, csvSeparator)
	return lineData[0]
}

func GetStringSliceFromModel(model interface{}) []string {
	var line []string
	v := reflect.ValueOf(model)
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i).Interface()
		line = append(line, fmt.Sprint(fieldValue))
	}
	return line
}
