package test

import "database/sql"
import _ "github.com/go-sql-driver/mysql" //MySQL driver used for instantiating the test database connection

//NewDBConnection returns a database instance to be used in tests
func NewDBConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(localhost)/")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
