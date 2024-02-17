// api_faculty.go
package api

import (
	"net/http"
	"github.com/RabbitLuke/seminar-search/dbQueries"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CreateUserRequest struct {
	Name string `json:"name"`
}

type DeleteUserRequest struct {
	FacultyID int `json:"facultyID"`
}

type UpdateUserRequest struct {
    FacultyID int    `json:"facultyID"`
    Name      string `json:"name"`
}

func CreateHandler(c *gin.Context) {
	var reqBody CreateUserRequest
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := query.InsertFaculty(reqBody.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func DeleteHandler(c *gin.Context) {
	var reqBody DeleteUserRequest
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := query.DeleteFaculty(reqBody.FacultyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func UpdateHandler(c *gin.Context) {
    var reqBody UpdateUserRequest
    if err := c.BindJSON(&reqBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := query.UpdateFaculty(reqBody.FacultyID, reqBody.Name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusOK)
}

func SelectHandler(c *gin.Context) {
    users, err := query.SelectFaculties()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, users)
}

func SelectByIDHandler(c *gin.Context) {
    facultyIDStr := c.Param("facultyID")
    facultyID, err := strconv.Atoi(facultyIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid facultyID"})
        return
    }

    user, err := query.SelectFacultyByID(facultyID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}
