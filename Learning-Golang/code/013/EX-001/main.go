package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// Data Source Name (dsn): Describes details about the connecting to DB
func main() {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   "localhost:5432",
		User:   url.UserPassword("user", "password"),
		Path:   "dbname",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		fmt.Println("sql.Open", err)
		return
	}
	defer func() {
		_ = db.Close()
		fmt.Println("closed")
	}()

	// 1. Let's access DB from Go!!
	// Get value from DB
	row := db.QueryRowContext(context.Background(), "SELECT birth_year FROM users WHERE name = 'jinhee'")
	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRowContext", err)
		return
	}

	// Scan the value
	var birth_year int64

	if err := row.Scan(&birth_year); err != nil {
		fmt.Println("row.Scan", err)
		return
	}

	fmt.Println("birth_year", birth_year)

	// Display name & bday from all users
	rows, err := db.QueryContext(context.Background(), "SELECT name, birth_year FROM users")
	if err != nil {
		fmt.Println("row.Scan", err)
		return
	}
	defer func() {
		_ = rows.Close()
	}()

	if rows.Err() != nil {
		fmt.Println("row.Err()", err)
		return
	}

	for rows.Next() {
		var name string
		var birth_year int64

		if err := rows.Scan(&name, &birth_year); err != nil {
			fmt.Println("rows.Scan", err)
			return
		}
		fmt.Println("name", name, "birth_year", birth_year)
	}

}
