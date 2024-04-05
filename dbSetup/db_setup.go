
package dbSetup

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
    connectionString := "dbStringHere"
    var err error
    DB, err = sql.Open("mysql", connectionString)
    if err != nil {
        return err
    }

    err = DB.Ping()
    if err != nil {
        return err
    }

    fmt.Println("Connected to the database")

    return nil
}

func CloseDB() {
    if DB != nil {
        DB.Close()
        fmt.Println("Closed the database connection")
    }
}
