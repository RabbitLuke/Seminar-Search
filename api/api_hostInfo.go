package api 

import (
	"net/http"
	"github.com/RabbitLuke/seminar-search/dbQueries"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CreateAHostRequest struct {
	HostID              int    `json:"HostID"`
	F_name              string `json:"First_Name"`
	L_name              string `json:"Last_Name"`
	Faculty             int    `json:"Faculty"`
	Qualifications int `json:"Qualifications"`
	Years_of_Experience int    `json:"Years_of_Experience"`
	Email               string `json:"Email"`
	Profile_pic         string `json:"Profile_pic"`
	Password            string `json:"Password"`
}

type DeleteAHostRequest struct {
	HostID int `json:"HostID"`
}

type UpdateAHostRequest struct {
	HostID              int    `json:"HostID"`
	F_name              string `json:"First_Name"`
	L_name              string `json:"Last_Name"`
	Faculty             int    `json:"Faculty"`
	Qualifications int `json:"Qualifications"`
	Years_of_Experience int    `json:"Years_of_Experience"`
	Email               string `json:"Email"`
	Profile_pic         string `json:"Profile_pic"`
	Password            string `json:"Password"`
}

func CreateHostHandler(c *gin.Context) {
	var reqBody CreateAHostRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := query.InsertHost(reqBody.F_name, reqBody.L_name, reqBody.Faculty, reqBody.Qualifications, reqBody.Years_of_Experience, reqBody.Email, reqBody.Profile_pic, reqBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func DeleteHostHandler(c *gin.Context) {
    
    HostID, err := strconv.Atoi(c.Param("HostID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid HostID"})
        return
    }

    
    err = query.DeleteHost(HostID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

   
    c.Status(http.StatusOK)
}

func UpdateHostHandler(c *gin.Context) {
	var reqBody UpdateAHostRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := query.UpdateHost(reqBody.HostID, reqBody.F_name, reqBody.L_name, reqBody.Faculty, reqBody.Qualifications, reqBody.Years_of_Experience, reqBody.Email, reqBody.Profile_pic, reqBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func SelectHostByIDHandler(c *gin.Context) {
	hostIDStr := c.Param("HostID")
	hostID, err := strconv.Atoi(hostIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hostID"})
		return
	}

	user, err := query.SelectHostByID(hostID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}