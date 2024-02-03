// main.go
package main

import (
	"fmt"
    dbSetup "github.com/RabbitLuke/seminar-search/dbSetup"
)

func main() {
    err := dbSetup.InitDB()
    if err != nil {
        fmt.Println("Error initializing the database:", err)
        return
    }
    defer dbSetup.CloseDB()
    

    // Your application logic here
    // You can use the 'db' variable from dbSetup package to perform database operations
}
