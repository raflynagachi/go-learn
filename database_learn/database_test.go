package databaselearn

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "nagachi@mail.co.id"
	comment := "Selamat datang yagesya"

	query := "INSERT INTO comment (email, comment) values(?,?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("success insert new comment with insertId", insertId)
}

func TestPrepStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comment (email, comment) VALUES (?,?)"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "nagachi" + strconv.Itoa(i) + "@mail.com"
		comment := "comment" + strconv.Itoa(i)

		res, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("comment id:", id)

	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// do transaction
	query := "INSERT INTO comment (email, comment) VALUES (?,?)"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "nagachi" + strconv.Itoa(i) + "@mail.com"
		comment := "comment" + strconv.Itoa(i)

		res, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("comment id:", id)

	}

	// err = tx.Commit()
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
