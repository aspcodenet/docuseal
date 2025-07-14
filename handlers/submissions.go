package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/repo/models"
	"gorm.io/gorm"
)

func CreateSubmission(c *gin.Context) {
	templateID, _ := strconv.Atoi(c.Param("id"))

	// Parse form data
	if err := c.Request.ParseForm(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	values := make(map[string]string)
	for key, value := range c.Request.PostForm {
		if len(value) > 0 {
			values[key] = value[0]
		}
	}

	valuesJSON, _ := json.Marshal(values)

	email, _ := c.Get("email")
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	db.Where("email = ?", email).First(&user)

	submission := models.Submission{
		TemplateID: uint(templateID),
		UserID:     user.ID,
		Values:     string(valuesJSON),
	}
	db.Create(&submission)

	c.Redirect(http.StatusFound, "/submissions/"+strconv.Itoa(int(submission.ID)))
}

func GetSubmission(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var submission models.Submission
	db.Preload("Template").First(&submission, id)

	var values map[string]string
	json.Unmarshal([]byte(submission.Values), &values)

	c.HTML(http.StatusOK, "submission.html", gin.H{
		"Submission": submission,
		"Values":     values,
	})
}

func NewSubmission(c *gin.Context) {
	templateID := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var template models.Template
	db.Preload("Fields").First(&template, templateID)

	c.HTML(http.StatusOK, "new_submission.html", gin.H{
		"Template": template,
	})
}
