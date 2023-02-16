package MainFlow

import (
	"aldaron/Constant"
	ExecuteQueryFlow "aldaron/Flow/ExecuteQuery"
	FormatQuery "aldaron/Flow/FormatQuery"
	ReadDataFlow "aldaron/Flow/ReadData"
	DbConfigModel "aldaron/Model/DbConfig"
	QueryConfigModel "aldaron/Model/QueryConfig"
	"aldaron/Utility/Database"
	"aldaron/Utility/Helper"
)

func getListTableSchema() []string {
	listTableSchemaContent := Helper.ReadCsvFile(Constant.INPUT_ORDER_DIR)
	var listTableSchema []string
	for _, tableSchema := range listTableSchemaContent {
		listTableSchema = append(listTableSchema, tableSchema[0])
	}
	return listTableSchema
}

func Main() {
	Helper.CleanOutputDirectory()
	var dbConfig DbConfigModel.DbConfig = DbConfigModel.DbConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "",
		DbName:   "postgres",
		SslMode:  "disable",
	}
	db := Database.GetDbConnection(dbConfig)
	var queryConfig QueryConfigModel.QueryConfig
	listTableSchema := getListTableSchema()
	for _, tableSchema := range listTableSchema {
		queryConfig = ReadDataFlow.Main(tableSchema)
		queryConfig = FormatQuery.Main(queryConfig)
		ExecuteQueryFlow.Main(db, queryConfig)
	}
	Database.CloseDbConnection(db)
}
