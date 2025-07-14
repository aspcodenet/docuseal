package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/repo/models"
	"gorm.io/gorm"
)

func APICreateSubmission(c *gin.Context) {
	var submissionData struct {
		TemplateID uint              `json:"template_id"`
		Values     map[string]string `json:"values"`
	}

	if err := c.ShouldBindJSON(&submissionData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valuesJSON, _ := json.Marshal(submissionData.Values)

	user := c.MustGet("user").(models.User)
	db := c.MustGet("db").(*gorm.DB)

	submission := models.Submission{
		TemplateID: submissionData.TemplateID,
		UserID:     user.ID,
		Values:     string(valuesJSON),
	}
	db.Create(&submission)

	c.JSON(http.StatusOK, submission)
}

func APIGetSubmissions(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	db := c.MustGet("db").(*gorm.DB)

	var submissions []models.Submission
	db.Where("user_id = ?", user.ID).Find(&submissions)

	c.JSON(http.StatusOK, submissions)
}

func APIGetSubmission(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var submission models.Submission
	db.First(&submission, id)

	var values map[string]string
	json.Unmarshal([]byte(submission.Values), &values)

	c.JSON(http.StatusOK, gin.H{
		"submission": submission,
		"values":     values,
	})
}
