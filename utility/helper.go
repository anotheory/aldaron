package utility

import (
	"aldaron/constant"
	"encoding/csv"
	"fmt"
	"os"
)

type HelperUtility struct{}

func (helper HelperUtility) CheckArrayContainString(arrayString []string, findString string) bool {
	for _, str := range arrayString {
		if findString == str {
			return true
		}
	}
	return false
}

func (helper HelperUtility) ConvertStringPointerArrayToValue(input [][]*string) [][]string {
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

func (helper HelperUtility) ValidateExistFile(filePath string) {
	if _, err := os.Stat(filePath); err != nil {
		panic(fmt.Sprintf("[%s] file doesn't exist!", filePath))
	}
}

func (helper HelperUtility) CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func (helper HelperUtility) ReadCsvFile(filePath string) [][]string {
	dataFile, err := os.Open(filePath)
	helper.CheckError(err)
	defer dataFile.Close()
	dataContent, err := csv.NewReader(dataFile).ReadAll()
	helper.CheckError(err)
	return dataContent
}

func (helper HelperUtility) ReadFileAsRawContent(filePath string) string {
	rawContent, err := os.ReadFile(filePath)
	helper.CheckError(err)
	return string(rawContent)
}

func (this HelperUtility) CleanOutputDirectory() {
	err := os.RemoveAll(constant.OUTPUT_DIR)
	this.CheckError(err)
	err = os.Mkdir(constant.OUTPUT_DIR, 0777)
	this.CheckError(err)
}
