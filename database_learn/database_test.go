package databaselearn

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {
	db, err := sql.Open("mysql", "adminer:password@tcp(localhost:3306)/godata")
	defer db.Close()

	if err != nil {
		panic(err)
	}
}
