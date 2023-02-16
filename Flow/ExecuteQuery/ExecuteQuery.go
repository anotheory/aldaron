package ExecuteQuery

import (
	"aldaron/Constant"
	QueryConfigModel "aldaron/Model/QueryConfig"
	"aldaron/Utility/Database"
	"aldaron/Utility/Helper"
	"database/sql"
	"encoding/csv"
	"os"
	"strings"
)

func formatQueryOutput(output *sql.Rows) ([]string, [][]string) {
	columns, _ := output.Columns()
	columnList := make([]*string, len(columns))
	for i := range columns {
		columnList[i] = &columns[i]
	}
	var results [][]*string
	for output.Next() {
		strs := make([]*string, len(columns))
		vals := make([]interface{}, len(columns))
		for i := range vals {
			vals[i] = &strs[i]
		}
		if err := output.Scan(vals...); err != nil {
			panic(err)
		}
		results = append(results, strs)
	}
	resultArray := Helper.ConvertStringPointerArrayToValue(results)
	return columns, resultArray
}

func writeOutputToFile(tableSchema string, outputColumns []string, outputArray [][]string) {
	outputFilePath := strings.Join([]string{Constant.OUTPUT_DIR, tableSchema}, "/") + ".csv"
	outputFile, err := os.Create(outputFilePath)
	Helper.CheckError(err)
	outputWriter := csv.NewWriter(outputFile)
	// Write column header rows
	outputWriter.Write(outputColumns)
	// Iterate through each records to write to output CSV file
	outputWriter.WriteAll(outputArray)
	outputFile.Close()
}

func Main(db *sql.DB, queryConfig QueryConfigModel.QueryConfig) {
	rows := Database.ExecuteQuery(db, queryConfig.Input.QueryString)
	outputColumns, outputArray := formatQueryOutput(rows)
	writeOutputToFile(queryConfig.TableSchema, outputColumns, outputArray)
}
