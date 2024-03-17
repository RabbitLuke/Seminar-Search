// main.go
package main

import (
	"log"
	"net/http"

	"github.com/RabbitLuke/seminar-search/api"
	dbSetup "github.com/RabbitLuke/seminar-search/dbSetup"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors"
	auth "github.com/RabbitLuke/seminar-search/auth"
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
		apiFaculty.DELETE("/remove/:facultyID", api.DeleteHandler)
		apiFaculty.PUT("/update", api.UpdateHandler)
		apiFaculty.GET("/all", api.SelectHandler)
		apiFaculty.GET("/distinct/:facultyID", api.SelectByIDHandler)
	}

	apiSeminar := router.Group("/seminar")
	{
		apiSeminar.POST("/create", api.CreateSeminarHandler)
		apiSeminar.DELETE("/remove/:seminarID", api.DeleteSeminarHandler)
		apiSeminar.PUT("/update", api.UpdateSeminarHandler)
		apiSeminar.GET("/all", api.SelectSeminarHandler)
		apiSeminar.GET("/distinct/:seminarID", api.SelectSeminarByIDHandler)
	}

	apiUser := router.Group("/user")
	{
		apiUser.POST("/create", api.CreateUserHandler)
		apiUser.DELETE("/remove/:UserID", api.DeleteUserHandler)
		apiUser.PUT("/update/", api.UpdateUserHandler)
		apiUser.GET("/distinct/:UserID", api.SelectUserByIDHandler)
	}
	
	apiHost := router.Group("/host")
	{
		apiHost.POST("/create", api.CreateHostHandler)
		apiHost.DELETE("/remove/:HostID", api.DeleteHostHandler)
		apiHost.PUT("/update", api.UpdateHostHandler)
		apiHost.GET("/distinct/:HostID", api.SelectHostByIDHandler)
	}

	apiQualification := router.Group("/qual")
	{
		apiQualification.GET("/all", api.SelectQualificationHandler)
	}

	router.POST("/authenticate", auth.Authenticate)
	router.POST("/refresh-token", auth.RefreshToken)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
	

	//CORS stuff
	//router.Use(gin.WrapH(cors.Default().Handler(http.DefaultServeMux)))

	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"127.0.0.1:8000"}, // Replace with your frontend's URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	})

	router.Use(gin.WrapH(corsConfig.Handler(http.DefaultServeMux)))
}
