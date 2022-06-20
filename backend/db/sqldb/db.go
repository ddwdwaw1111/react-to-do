package sqldb

import (
	"database/sql"
)

// ConnectDB opens a connection to the database
func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/my_db")

	if err != nil {
		panic(err.Error())
	}

	return db
}
