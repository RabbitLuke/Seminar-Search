package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		splitToken := strings.Split(authorizationHeader, "Bearer ")
		if len(splitToken) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := splitToken[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("test"), nil
		})

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		exp, ok := claims["exp"].(float64)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		expTime := time.Unix(int64(exp), 0)
		if time.Now().After(expTime) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
