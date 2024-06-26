// main.go
package main

import (
	"log"
	"net/http"
	"github.com/RabbitLuke/seminar-search/api"
	auth "github.com/RabbitLuke/seminar-search/auth"
	dbSetup "github.com/RabbitLuke/seminar-search/dbSetup"
	middleware "github.com/RabbitLuke/seminar-search/middleware"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors"
)

func main() {

	err := dbSetup.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbSetup.CloseDB()

	router := gin.Default()

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
		apiSeminar.GET("/getseminarbyfaculty/:facultyId", api.GetSeminarsByFaculty)
	}

	apiUser := router.Group("/user")
	{
		apiUser.POST("/create", api.CreateUserHandler)
		apiUser.DELETE("/remove/:UserID", api.DeleteUserHandler)
		apiUser.PUT("/update/", api.UpdateUserHandler)
		apiUser.GET("/distinct/:UserID", api.SelectUserByIDHandler)
		apiUser.GET("/getuser", middleware.JWTMiddleware(), api.GetUserByEmail)
		apiUser.GET("/getuserdashboard", middleware.JWTMiddleware(), api.GetUserFacultyAndSeminars)
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

	
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"FrontendURL"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	})

	router.Use(gin.WrapH(corsConfig.Handler(http.DefaultServeMux)))
}
