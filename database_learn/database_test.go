package databaselearn

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {
	db, err := sql.Open("mysql", "adminer:password@tcp(localhost:3306)/godata")
	defer db.Close()

	if err != nil {
		panic(err)
	}
}

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, name) values ('go', 'Nagachi')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT id,name FROM customer"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, ":", name)
	}
}

func TestQuerySqlAdvanced(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date sql.NullTime
		var created_at time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("==============")
		fmt.Println(id, name, email.String, balance, rating, birth_date.Time, married, created_at)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// username will pass if username="admin'; #""
	username := "admin"
	password := "admin"

	query := "SELECT username FROM user WHERE username='" + username +
		"' AND password='" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var username string
		rows.Scan(&username)
		fmt.Println(username, "logged in")
	} else {
		fmt.Println("failed to login")
	}
}

func TestSQLWithParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// username will NOT PASS if username="admin'; #""
	var username = "admin"
	var password = "admin"

	query := "SELECT username FROM user WHERE username=? AND password=? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var username string
		rows.Scan(&username)
		fmt.Println(username, "has logged in successfully")
	} else {
		fmt.Println("Failed to login")
	}
}
