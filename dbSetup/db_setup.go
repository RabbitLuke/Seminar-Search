// db_setup.go
package dbSetup

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB() error {
    connectionString := "root:Banana10#@tcp(localhost:3306)/seminarsearch"
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

// CloseDB closes the database connection
func CloseDB() {
    if DB != nil {
        DB.Close()
        fmt.Println("Closed the database connection")
    }
}
