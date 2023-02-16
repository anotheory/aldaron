package ReadData

import (
	"aldaron/Constant"
	ReadDataLogic "aldaron/Logic/ReadData"
	QueryConfigModel "aldaron/Model/QueryConfig"
	"aldaron/Utility/Helper"
	"fmt"
)

func validateExistedRequiredInput(tableSchema string) QueryConfigModel.QueryConfig {
	columnFilePath := fmt.Sprintf("%s/%s/%s", Constant.INPUT_DIR, tableSchema, Constant.COLUMN_FILE_NAME)
	Helper.ValidateExistFile(columnFilePath)
	dataFilePath := fmt.Sprintf("%s/%s/%s", Constant.INPUT_DIR, tableSchema, Constant.DATA_FILE_NAME)
	Helper.ValidateExistFile(dataFilePath)
	queryFilePath := fmt.Sprintf("%s/%s/%s", Constant.INPUT_DIR, tableSchema, Constant.QUERY_FILE_NAME)
	Helper.ValidateExistFile(queryFilePath)
	valueQueryFilePath := fmt.Sprintf("%s/%s/%s", Constant.INPUT_DIR, tableSchema, Constant.VALUE_QUERY_FILE_NAME)
	Helper.ValidateExistFile(valueQueryFilePath)
	queryConfig := QueryConfigModel.QueryConfig{
		TableSchema: tableSchema,
		Input: QueryConfigModel.InputData{
			FilePath: QueryConfigModel.InputFilePath{
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
	return Helper.ReadCsvFile(columnFilePath)[0]
}

func readDataContent(dataFilePath string) [][]string {
	return Helper.ReadCsvFile(dataFilePath)
}

func readQueryContent(queryFilePath string) string {
	queryString := Helper.ReadFileAsRawContent(queryFilePath)
	ReadDataLogic.ValidateQueryPlaceholder(queryString)
	return queryString
}

func readValueQueryContent(valueQueryFilePath string) string {
	valueQueryString := Helper.ReadFileAsRawContent(valueQueryFilePath)
	ReadDataLogic.ValidateValueQueryPlaceholder(valueQueryString)
	return valueQueryString
}

func Main(tableSchema string) QueryConfigModel.QueryConfig {
	queryConfig := validateExistedRequiredInput(tableSchema)
	queryConfig.Input.DataColumn = readColumnContent(queryConfig.Input.FilePath.Column)
	queryConfig.Input.DataGrid = readDataContent(queryConfig.Input.FilePath.Data)
	queryConfig.Input.QueryString = readQueryContent(queryConfig.Input.FilePath.Query)
	queryConfig.Input.ValueQueryString = readValueQueryContent(queryConfig.Input.FilePath.ValueQuery)
	return queryConfig
}
