package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func connectToMariaDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Banana10#@tcp(localhost:3306)/mydb")
	if err != nil {
		return nil, err
	}

	// Ping the MariaDB server to ensure connectivity
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MariaDB!")
	return db, nil
}
