package api

import (
	"net/http"
	"strconv"

	query "github.com/RabbitLuke/seminar-search/dbQueries"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CreateAUserRequest struct {
	UserId      string `json:"UserID"`
	F_name      string `json:"First_Name"`
	L_name      string `json:"Last_Name"`
	Faculty     int    `json:"Faculty"`
	Email       string `json:"Email"`
	Profile_pic string `json:"Profile_pic"`
	Password    string `json:"Password"`
}

type DeleteAUserRequest struct {
	UserID int `json:"UserID"`
}

type UpdateAUserRequest struct {
	UserId      int    `json:"UserID"`
	F_name      string `json:"First_Name"`
	L_name      string `json:"Last_Name"`
	Faculty     int    `json:"Faculty"`
	Email       string `json:"Email"`
	Profile_pic string `json:"Profile_pic"`
	Password    string `json:"Password"`
}

func CreateUserHandler(c *gin.Context) {
	var reqBody CreateAUserRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := query.InsertUser(reqBody.F_name, reqBody.L_name, reqBody.Faculty, reqBody.Email, reqBody.Profile_pic, reqBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func DeleteUserHandler(c *gin.Context) {
	// Extract UserID from the URL parameter
	UserID, err := strconv.Atoi(c.Param("UserID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID"})
		return
	}

	// Call the DeleteUser function with the extracted UserID
	err = query.DeleteUser(UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success status
	c.Status(http.StatusOK)
}

func UpdateUserHandler(c *gin.Context) {
	var reqBody UpdateAUserRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := query.UpdateUser(reqBody.UserId, reqBody.F_name, reqBody.L_name, reqBody.Faculty, reqBody.Email, reqBody.Profile_pic, reqBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func SelectUserByIDHandler(c *gin.Context) {
	userIDStr := c.Param("UserID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID"})
		return
	}

	user, err := query.SelectUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserFacultyAndSeminars(c *gin.Context) {
	//extract the user from the jwt
	claims := c.Value("user").(jwt.MapClaims)
	email := claims["Email"].(string)
	data, err := query.GetUserFacultyAndSeminary(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetUserByEmail(c *gin.Context) {
	//extract the user from the jwt
	claims := c.Value("user").(jwt.MapClaims)
	email := claims["Email"].(string)
	data, err := query.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("#######################")
		c.Next()
	}
}
