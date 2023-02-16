package FormatQuery

import (
	"aldaron/Constant"
	QueryConfigModel "aldaron/Model/QueryConfig"
	"fmt"
	"strings"
)

func replaceColumnString(queryString string, dataColumn []string) string {
	queryString = strings.Replace(queryString, Constant.PLACEHOLDER_COLUMNS, fmt.Sprintf("(%s)", strings.Join(dataColumn, ",")), 1)
	return queryString
}

func replaceValueString(queryString string, valueQueryString string, dataGrid [][]string) string {
	var valueStringArr []string
	var rowString string
	for rowIdx, _ := range dataGrid {
		rowString = valueQueryString
		for colIdx, _ := range dataGrid[rowIdx] {
			rowString = strings.Replace(rowString, fmt.Sprintf("[val:%d]", colIdx+1), dataGrid[rowIdx][colIdx], -1)
		}
		valueStringArr = append(valueStringArr, fmt.Sprintf("(%s)", rowString))
	}
	queryString = strings.Replace(queryString, Constant.PLACEHOLDER_VALUES, strings.Join(valueStringArr, ","), 1)
	return queryString
}

func Main(queryConfig QueryConfigModel.QueryConfig) QueryConfigModel.QueryConfig {
	queryConfig.Input.QueryString = replaceColumnString(queryConfig.Input.QueryString, queryConfig.Input.DataColumn)
	queryConfig.Input.QueryString = replaceValueString(queryConfig.Input.QueryString, queryConfig.Input.ValueQueryString, queryConfig.Input.DataGrid)
	return queryConfig
}
