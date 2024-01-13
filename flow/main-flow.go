package flow

import (
	"aldaron/constant"
	"aldaron/models/config"
	"aldaron/utility"
)

type MainFlow struct {
	readDataFlow     ReadDataFlow
	formatQueryFlow  FormatQueryFlow
	executeQueryFlow ExecuteQueryFlow
	helper           utility.HelperUtility
	database         utility.DatabaseUtility
}

func (flow MainFlow) getListTableSchema() []string {
	listTableSchemaContent := flow.helper.ReadCsvFile(constant.INPUT_ORDER_DIR)
	var listTableSchema []string
	for _, tableSchema := range listTableSchemaContent {
		listTableSchema = append(listTableSchema, tableSchema[0])
	}
	return listTableSchema
}

func (flow MainFlow) Main() {
	flow.helper.CleanOutputDirectory()
	var dbConfig config.DbConfig = config.DbConfig{
		Host:     "localhost",
		Port:     5082,
		User:     "root",
		Password: "",
		DbName:   "ngernturbo",
		SslMode:  "disable",
	}
	db := flow.database.GetDbConnection(dbConfig)
	var queryConfig config.QueryConfig
	listTableSchema := flow.getListTableSchema()
	for _, tableSchema := range listTableSchema {
		queryConfig = flow.readDataFlow.Main(tableSchema)
		queryConfig = flow.formatQueryFlow.Main(queryConfig)
		flow.executeQueryFlow.Main(db, queryConfig)
	}
	flow.database.CloseDbConnection(db)
}
