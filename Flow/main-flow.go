package flow

import (
	"aldaron/constant"
	executeQueryFlow "aldaron/flow/execute-query"
	formatQueryFlow "aldaron/flow/format-query"
	readDataFlow "aldaron/flow/read-data"
	"aldaron/model/config"
	"aldaron/utility/database"
	"aldaron/utility/helper"
)

func getListTableSchema() []string {
	listTableSchemaContent := helper.ReadCsvFile(constant.INPUT_ORDER_DIR)
	var listTableSchema []string
	for _, tableSchema := range listTableSchemaContent {
		listTableSchema = append(listTableSchema, tableSchema[0])
	}
	return listTableSchema
}

func Main() {
	helper.CleanOutputDirectory()
	var dbConfig config.DbConfig = config.DbConfig{
		Host:     "localhost",
		Port:     5082,
		User:     "root",
		Password: "",
		DbName:   "ngernturbo",
		SslMode:  "disable",
	}
	db := database.GetDbConnection(dbConfig)
	var queryConfig config.QueryConfig
	listTableSchema := getListTableSchema()
	for _, tableSchema := range listTableSchema {
		queryConfig = readDataFlow.Main(tableSchema)
		queryConfig = formatQueryFlow.Main(queryConfig)
		executeQueryFlow.Main(db, queryConfig)
	}
	database.CloseDbConnection(db)
}
