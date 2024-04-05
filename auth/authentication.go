package auth

import (
	"net/http"
	"time"
	"github.com/RabbitLuke/seminar-search/dbSetup"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	UserId      int    `json:"UserID"`
	F_name      string `json:"First_Name"`
	L_name      string `json:"Last_Name"`
	Faculty     int    `json:"Faculty"`
	Email       string `json:"Email"`
	Profile_pic string `json:"Profile_pic"`
	Password    string `json:"Password"`
}

type TokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func Authenticate(c *gin.Context) {
	var user UserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if dbSetup.DB == nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	var userCount int
	userErr := dbSetup.DB.QueryRow("SELECT COUNT(*) FROM user_information WHERE Email = ? and Password = ?", user.Email, user.Password).Scan(&userCount)
	if userErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		return
	}

	var hostCount int
	hostErr := dbSetup.DB.QueryRow("SELECT COUNT(*) FROM host_information WHERE Email = ? and Password = ?", user.Email, user.Password).Scan(&hostCount)
	if hostErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		return
	}

	if userCount <= 0 && hostCount <= 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email":  user.Email,
		"IsHost": hostCount > 0,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte("test"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte("test"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := TokenResponse{
		Token:        tokenString,
		RefreshToken: refreshTokenString,
	}

	c.JSON(http.StatusOK, response)
}

func RefreshToken(c *gin.Context) {
	var refreshRequest struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&refreshRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.Parse(refreshRequest.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("test"), nil 
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	username, ok := claims["username"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	})


	newTokenString, err := newToken.SignedString([]byte("test"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	newRefreshTokenString, err := refreshToken.SignedString([]byte("test"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := TokenResponse{
		Token:        newTokenString,
		RefreshToken: newRefreshTokenString,
	}

	c.JSON(http.StatusOK, response)
}
