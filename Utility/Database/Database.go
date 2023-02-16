package Database

import (
	DbConfigModel "aldaron/Model/DbConfig"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetDbConnection(dbConfig DbConfigModel.DbConfig) *sql.DB {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DbName,
		dbConfig.SslMode,
	)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}

func ExecuteQuery(db *sql.DB, queryString string) *sql.Rows {
	resultRows, err := db.Query(queryString)
	if err != nil {
		panic(err)
	}
	return resultRows
}

func CloseDbConnection(db *sql.DB) {
	db.Close()
}
