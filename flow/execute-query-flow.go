package flow

import (
	"aldaron/constant"
	"aldaron/models/config"
	"aldaron/utility"
	"database/sql"
	"encoding/csv"
	"os"
	"strings"
)

type ExecuteQueryFlow struct {
	helper   utility.HelperUtility
	database utility.DatabaseUtility
}

func (flow ExecuteQueryFlow) formatQueryOutput(output *sql.Rows) ([]string, [][]string) {
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
	resultArray := flow.helper.ConvertStringPointerArrayToValue(results)
	return columns, resultArray
}

func (flow ExecuteQueryFlow) writeOutputToFile(tableSchema string, outputColumns []string, outputArray [][]string) {
	outputFilePath := strings.Join([]string{constant.OUTPUT_DIR, tableSchema}, "/") + ".csv"
	outputFile, err := os.Create(outputFilePath)
	flow.helper.CheckError(err)
	outputWriter := csv.NewWriter(outputFile)
	// Write column header rows
	outputWriter.Write(outputColumns)
	// Iterate through each records to write to output CSV file
	outputWriter.WriteAll(outputArray)
	outputFile.Close()
}

func (flow ExecuteQueryFlow) Main(db *sql.DB, queryConfig config.QueryConfig) {
	rows := flow.database.ExecuteQuery(db, queryConfig.Input.QueryString)
	outputColumns, outputArray := flow.formatQueryOutput(rows)
	flow.writeOutputToFile(queryConfig.TableSchema, outputColumns, outputArray)
}
