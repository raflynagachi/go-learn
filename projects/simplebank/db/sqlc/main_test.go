package db_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	db "github.com/raflynagachi/simplebank/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/simplebank_db?sslmode=disable"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = db.New(conn)
	os.Exit(m.Run())
}
