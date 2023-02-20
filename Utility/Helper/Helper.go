package Helper

import (
	"aldaron/Constant"
	"encoding/csv"
	"fmt"
	"os"
)

func CheckArrayContainString(arrayString []string, findString string) bool {
	for _, str := range arrayString {
		if findString == str {
			return true
		}
	}
	return false
}

func ConvertStringPointerArrayToValue(input [][]*string) [][]string {
	var output [][]string
	for rowIdx, _ := range input {
		output = append(output, []string{})
		for colIdx, _ := range input[rowIdx] {
			if input[rowIdx][colIdx] == nil {
				output[rowIdx] = append(output[rowIdx], "")
			} else {
				output[rowIdx] = append(output[rowIdx], *input[rowIdx][colIdx])
			}
		}
	}
	return output
}

func ValidateExistFile(filePath string) {
	if _, err := os.Stat(filePath); err != nil {
		panic(fmt.Sprintf("[%s] file doesn't exist!", filePath))
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadCsvFile(filePath string) [][]string {
	dataFile, err := os.Open(filePath)
	CheckError(err)
	defer dataFile.Close()
	dataContent, err := csv.NewReader(dataFile).ReadAll()
	CheckError(err)
	return dataContent
}

func ReadFileAsRawContent(filePath string) string {
	rawContent, err := os.ReadFile(filePath)
	CheckError(err)
	return string(rawContent)
}

func CleanOutputDirectory() {
	err := os.RemoveAll(Constant.OUTPUT_DIR)
	CheckError(err)
	err = os.Mkdir(Constant.OUTPUT_DIR, 0777)
	CheckError(err)
}
