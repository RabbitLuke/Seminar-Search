// main.go
package main

import (
    "log"
    //queryUsers "github.com/RabbitLuke/seminar-search/dbQueries"
    dbSetup "github.com/RabbitLuke/seminar-search/dbSetup"
    "github.com/gin-gonic/gin"
    "github.com/RabbitLuke/seminar-search/api"
)

func main() {
    
    err := dbSetup.InitDB()
    if err != nil {
        log.Fatal(err)
    }
    defer dbSetup.CloseDB()

    
	// name := "Archeology"

	// err = queryUsers.InsertUser(name)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
    // Delete a user with facultyID 1
    // err = queryUsers.DeleteUser(1)
    // if err != nil {
    //     log.Fatal(err)
    // } 
    // Your application logic here
    // You can use the 'db' variable from dbSetup package to perform database operations
    router := gin.Default()

	// Import and use functions from api_faculty.go
    apiGroup := router.Group("/api")
	{
		apiGroup.POST("/user", api.CreateHandler)
		apiGroup.DELETE("/user", api.DeleteHandler)
		apiGroup.PUT("/user", api.UpdateHandler)
		apiGroup.GET("/users", api.SelectHandler)
		apiGroup.GET("/user/:facultyID", api.SelectByIDHandler)
	}
        
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
    
}
