// db_setup.go
package dbSetup

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDB initializes the database connection
func InitDB() error {
    connectionString := "root:Banana10#@tcp(localhost:3306)/seminarsearch"
    var err error
    db, err = sql.Open("mysql", connectionString)
    if err != nil {
        return err
    }

    err = db.Ping()
    if err != nil {
        return err
    }

    fmt.Println("Connected to the database")

    return nil
}

// CloseDB closes the database connection
func CloseDB() {
    if db != nil {
        db.Close()
        fmt.Println("Closed the database connection")
    }
}
