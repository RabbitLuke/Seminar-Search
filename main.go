// main.go
package main

import (
    "log"
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

    router := gin.Default()

	// Import and use functions from api_faculty.go
    apiGroup := router.Group("/faculty")
	{
		apiGroup.POST("/create", api.CreateHandler)
		apiGroup.DELETE("/remove", api.DeleteHandler)
		apiGroup.PUT("/update", api.UpdateHandler)
		apiGroup.GET("/all", api.SelectHandler)
		apiGroup.GET("/distinct/:facultyID", api.SelectByIDHandler)
	}
        
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
    
}
