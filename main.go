// main.go
package main

import (
    "log"
    queryUsers "github.com/RabbitLuke/seminar-search/dbQueries"
    dbSetup "github.com/RabbitLuke/seminar-search/dbSetup"
)

func main() {
    
    err := dbSetup.InitDB()
    if err != nil {
        log.Fatal(err)
    }
    defer dbSetup.CloseDB()

    
	name := "Geography"

	err = queryUsers.InsertUser(name)
	if err != nil {
		log.Fatal(err)
	}
	
    
    // Your application logic here
    // You can use the 'db' variable from dbSetup package to perform database operations
}
