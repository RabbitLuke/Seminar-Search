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

// Authenticate authenticates the user and generates JWT tokens
func Authenticate(c *gin.Context) {
	var user UserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the user credentials (you can replace this with your own authentication logic)

	if dbSetup.DB == nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	//is user login
	var userCount int
	userErr := dbSetup.DB.QueryRow("SELECT COUNT(*) FROM user_information WHERE Email = ? and Password = ?", user.Email, user.Password).Scan(&userCount)
	if userErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		return
	}

	//is host login
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

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email":  user.Email,
		"IsHost": hostCount > 0,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(), // Token expires in 15 minutes
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte("test")) // Replace "your-secret-key" with your own secret key
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(), // Refresh token expires in 7 days
	})

	// Sign the refresh token with the secret key
	refreshTokenString, err := refreshToken.SignedString([]byte("test")) // Replace "your-secret-key" with your own secret key
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the tokens as a response
	response := TokenResponse{
		Token:        tokenString,
		RefreshToken: refreshTokenString,
	}

	c.JSON(http.StatusOK, response)
}

// RefreshToken generates a new access token using the refresh token
func RefreshToken(c *gin.Context) {
	var refreshRequest struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&refreshRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.Parse(refreshRequest.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("test"), nil // Replace "your-secret-key" with your own refresh secret key
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

	// Retrieve the username from the refresh token claims
	username, ok := claims["username"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Create a new access token for the user
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 15).Unix(), // Token expires in 15 minutes
	})

	// Sign the new token with the secret key
	newTokenString, err := newToken.SignedString([]byte("test")) // Replace "your-secret-key" with your own secret key
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Create refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), // Refresh token expires in 7 days
	})

	// Sign the refresh token with the secret key
	newRefreshTokenString, err := refreshToken.SignedString([]byte("test")) // Replace "your-secret-key" with your own secret key
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the new access token as a response
	response := TokenResponse{
		Token:        newTokenString,
		RefreshToken: newRefreshTokenString,
	}

	c.JSON(http.StatusOK, response)
}
