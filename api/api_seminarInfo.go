package api

import (
	"net/http"
	"strconv"
	query "github.com/RabbitLuke/seminar-search/dbQueries"
	"github.com/gin-gonic/gin"
)

type CreateSeminarRequest struct {
	Title      string  `json:"Title"`
	FacultyID  int     `json:"facultyID"`
	Duration   float32 `json:"Duration"`
	Date       string  `json:"Date"`
	Time       string  `json:"Time"`
	Location   string  `json:"Location"`
	NoOfSeats  int     `json:"no_of_seats"`
	CoverPhoto string  `json:"cover_photo"`
}

type DeleteSeminarRequest struct {
	SeminarID int `json:"seminarID"`
}

type UpdateSeminarRequest struct {
	SeminarID  int     `json:"seminarID"`
	Name       string  `json:"name"`
	Title      string  `json:"Title"`
	FacultyID  int     `json:"facultyID"`
	Duration   float32 `json:"Duration"`
	Date       string  `json:"Date"`
	Time       string  `json:"Time"`
	Location   string  `json:"Location"`
	NoOfSeats  int     `json:"no_of_seats"`
	CoverPhoto string  `json:"cover_photo"`
}

func CreateSeminarHandler(c *gin.Context) {
	var reqBody CreateSeminarRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := query.InsertSeminar(reqBody.Title, reqBody.FacultyID, reqBody.Duration, reqBody.Date, reqBody.Time, reqBody.Location, reqBody.NoOfSeats, reqBody.CoverPhoto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func DeleteSeminarHandler(c *gin.Context) {
	seminarID, err := strconv.Atoi(c.Param("seminarID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seminarID"})
		return
	}

	err = query.DeleteSeminar(seminarID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func UpdateSeminarHandler(c *gin.Context) {
	var reqBody UpdateSeminarRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := query.UpdateSeminar(reqBody.SeminarID, reqBody.Title, reqBody.FacultyID, reqBody.Duration, reqBody.Date, reqBody.Time, reqBody.Location, reqBody.NoOfSeats, reqBody.CoverPhoto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func SelectSeminarHandler(c *gin.Context) {
	users, err := query.SelectSeminars()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func SelectSeminarByIDHandler(c *gin.Context) {
	seminarIDStr := c.Param("seminarID")
	seminarID, err := strconv.Atoi(seminarIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seminarID"})
		return
	}

	user, err := query.SelectSeminarByID(seminarID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetSeminarsByFaculty(c *gin.Context) {
	facultyIdSTR := c.Param("facultyId")
	facultyId, err := strconv.Atoi(facultyIdSTR)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid facultyId"})
		return
	}

	user, err := query.GetSeminarsByFaculty(facultyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
