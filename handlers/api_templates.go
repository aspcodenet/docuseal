package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/repo/models"
	"gorm.io/gorm"
)

func APIGetTemplates(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	db := c.MustGet("db").(*gorm.DB)

	var templates []models.Template
	db.Where("user_id = ?", user.ID).Find(&templates)

	c.JSON(http.StatusOK, templates)
}

func APIGetTemplate(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var template models.Template
	db.Preload("Fields").First(&template, id)

	c.JSON(http.StatusOK, template)
}

func APICreateTemplate(c *gin.Context) {
	var template models.Template
	if err := c.ShouldBindJSON(&template); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.MustGet("user").(models.User)
	db := c.MustGet("db").(*gorm.DB)

	template.UserID = user.ID
	db.Create(&template)

	for _, field := range template.Fields {
		field.TemplateID = template.ID
		db.Create(&field)
	}

	c.JSON(http.StatusOK, template)
}

func APIUpdateTemplate(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var template models.Template
	db.First(&template, id)

	if err := c.ShouldBindJSON(&template); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&template)

	c.JSON(http.StatusOK, template)
}

func APIDeleteTemplate(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var template models.Template
	db.First(&template, id)
	db.Delete(&template)

	c.JSON(http.StatusOK, gin.H{"message": "Template deleted successfully"})
}
