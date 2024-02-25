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
    apiFaculty := router.Group("/faculty")
	{
		apiFaculty.POST("/create", api.CreateHandler)
		apiFaculty.DELETE("/remove", api.DeleteHandler)
		apiFaculty.PUT("/update", api.UpdateHandler)
		apiFaculty.GET("/all", api.SelectHandler)
		apiFaculty.GET("/distinct/:facultyID", api.SelectByIDHandler)
	}
    
	apiSeminar := router.Group("/seminar")
	{
		apiSeminar.POST("/create", api.CreateSeminarHandler)
		apiSeminar.DELETE("/remove", api.DeleteSeminarHandler)
		apiSeminar.PUT("/update", api.UpdateSeminarHandler)
		apiSeminar.GET("/all", api.SelectSeminarHandler)
		apiSeminar.GET("/distinct/:seminarID", api.SelectSeminarByIDHandler)
	}

	router.POST("/authenticate", Authenticate)
	router.POST("/refresh-token", RefreshToken)
	
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
    
}
