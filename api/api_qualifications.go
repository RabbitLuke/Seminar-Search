package api

import (
	"net/http"
	"github.com/RabbitLuke/seminar-search/dbQueries"
	"github.com/gin-gonic/gin"

)

func SelectQualificationHandler(c *gin.Context) {
    users, err := query.SelectQualifications()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, users)
}