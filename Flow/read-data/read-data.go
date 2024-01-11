package read_data

import (
	"aldaron/constant"
	readDataLogic "aldaron/logic/read-data"
	"aldaron/model/config"
	"aldaron/utility/helper"
	"fmt"
)

func validateExistedRequiredInput(tableSchema string) config.QueryConfig {
	columnFilePath := fmt.Sprintf("%s/%s/%s", constant.INPUT_DIR, tableSchema, constant.COLUMN_FILE_NAME)
	helper.ValidateExistFile(columnFilePath)
	dataFilePath := fmt.Sprintf("%s/%s/%s", constant.INPUT_DIR, tableSchema, constant.DATA_FILE_NAME)
	helper.ValidateExistFile(dataFilePath)
	queryFilePath := fmt.Sprintf("%s/%s/%s", constant.INPUT_DIR, tableSchema, constant.QUERY_FILE_NAME)
	helper.ValidateExistFile(queryFilePath)
	valueQueryFilePath := fmt.Sprintf("%s/%s/%s", constant.INPUT_DIR, tableSchema, constant.VALUE_QUERY_FILE_NAME)
	helper.ValidateExistFile(valueQueryFilePath)
	queryConfig := config.QueryConfig{
		TableSchema: tableSchema,
		Input: config.InputData{
			FilePath: config.InputFilePath{
				Column:     columnFilePath,
				Data:       dataFilePath,
				Query:      queryFilePath,
				ValueQuery: valueQueryFilePath,
			},
		},
	}
	return queryConfig
}

func readColumnContent(columnFilePath string) []string {
	return helper.ReadCsvFile(columnFilePath)[0]
}

func readDataContent(dataFilePath string) [][]string {
	return helper.ReadCsvFile(dataFilePath)
}

func readQueryContent(queryFilePath string) string {
	queryString := helper.ReadFileAsRawContent(queryFilePath)
	readDataLogic.ValidateQueryPlaceholder(queryString)
	return queryString
}

func readValueQueryContent(valueQueryFilePath string) string {
	valueQueryString := helper.ReadFileAsRawContent(valueQueryFilePath)
	readDataLogic.ValidateValueQueryPlaceholder(valueQueryString)
	return valueQueryString
}

func Main(tableSchema string) config.QueryConfig {
	queryConfig := validateExistedRequiredInput(tableSchema)
	queryConfig.Input.DataColumn = readColumnContent(queryConfig.Input.FilePath.Column)
	queryConfig.Input.DataGrid = readDataContent(queryConfig.Input.FilePath.Data)
	queryConfig.Input.QueryString = readQueryContent(queryConfig.Input.FilePath.Query)
	queryConfig.Input.ValueQueryString = readValueQueryContent(queryConfig.Input.FilePath.ValueQuery)
	return queryConfig
}
