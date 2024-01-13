package flow

import (
	"aldaron/constant"
	"aldaron/logic"
	"aldaron/models/config"
	"aldaron/utility"
	"fmt"
)

type ReadDataFlow struct {
	helper        utility.HelperUtility
	readDataLogic logic.ReadDataLogic
}

func (flow ReadDataFlow) validateExistedRequiredInput(tableSchema string) config.QueryConfig {
	columnFilePath := fmt.Sprintf("%s/%s/%s", constant.INPUT_DIR, tableSchema, constant.COLUMN_FILE_NAME)
	flow.helper.ValidateExistFile(columnFilePath)
	dataFilePath := fmt.Sprintf("%s/%s/%s", constant.INPUT_DIR, tableSchema, constant.DATA_FILE_NAME)
	flow.helper.ValidateExistFile(dataFilePath)
	queryFilePath := fmt.Sprintf("%s/%s/%s", constant.INPUT_DIR, tableSchema, constant.QUERY_FILE_NAME)
	flow.helper.ValidateExistFile(queryFilePath)
	valueQueryFilePath := fmt.Sprintf("%s/%s/%s", constant.INPUT_DIR, tableSchema, constant.VALUE_QUERY_FILE_NAME)
	flow.helper.ValidateExistFile(valueQueryFilePath)
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

func (flow ReadDataFlow) readColumnContent(columnFilePath string) []string {
	return flow.helper.ReadCsvFile(columnFilePath)[0]
}

func (flow ReadDataFlow) readDataContent(dataFilePath string) [][]string {
	return flow.helper.ReadCsvFile(dataFilePath)
}

func (flow ReadDataFlow) readQueryContent(queryFilePath string) string {
	queryString := flow.helper.ReadFileAsRawContent(queryFilePath)
	flow.readDataLogic.ValidateQueryPlaceholder(queryString)
	return queryString
}

func (flow ReadDataFlow) readValueQueryContent(valueQueryFilePath string) string {
	valueQueryString := flow.helper.ReadFileAsRawContent(valueQueryFilePath)
	flow.readDataLogic.ValidateValueQueryPlaceholder(valueQueryString)
	return valueQueryString
}

func (flow ReadDataFlow) Main(tableSchema string) config.QueryConfig {
	queryConfig := flow.validateExistedRequiredInput(tableSchema)
	queryConfig.Input.DataColumn = flow.readColumnContent(queryConfig.Input.FilePath.Column)
	queryConfig.Input.DataGrid = flow.readDataContent(queryConfig.Input.FilePath.Data)
	queryConfig.Input.QueryString = flow.readQueryContent(queryConfig.Input.FilePath.Query)
	queryConfig.Input.ValueQueryString = flow.readValueQueryContent(queryConfig.Input.FilePath.ValueQuery)
	return queryConfig
}
