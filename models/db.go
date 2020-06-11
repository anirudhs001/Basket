package models

import "database/sql"

var db *sql.DB //not exported

func init() {

	var err error
	db, err = sql.Open("postgres", "postgres://temp_user:password@localhost/basket?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
}
