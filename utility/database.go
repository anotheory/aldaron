package utility

import (
	"aldaron/models/config"
	"database/sql"
	"fmt"
)

type DatabaseUtility struct{}

func (database DatabaseUtility) GetDbConnection(dbConfig config.DbConfig) *sql.DB {
	var connString string
	if dbConfig.Password == "" {
		connString = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.DbName, dbConfig.SslMode)
	} else {
		connString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DbName, dbConfig.SslMode)
	}
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}

func (database DatabaseUtility) ExecuteQuery(db *sql.DB, queryString string) *sql.Rows {
	resultRows, err := db.Query(queryString)
	if err != nil {
		panic(err)
	}
	return resultRows
}

func (database DatabaseUtility) CloseDbConnection(db *sql.DB) {
	db.Close()
}
