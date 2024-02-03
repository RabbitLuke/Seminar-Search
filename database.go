package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func connectDB() (*sql.DB, error) {
    db, err := sql.Open("seminarsearch", "root:banana10#@tcp(localhost:3306)/seminarsearch")
    if err != nil {
        return nil, err
    }
    return db, nil
}
