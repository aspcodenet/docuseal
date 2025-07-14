package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/repo/models"
	"gorm.io/gorm"
)

func CreateTemplate(c *gin.Context) {
	var template models.Template
	if err := c.ShouldBind(&template); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email, _ := c.Get("email")
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	db.Where("email = ?", email).First(&user)

	template.UserID = user.ID
	db.Create(&template)

	for _, field := range template.Fields {
		field.TemplateID = template.ID
		db.Create(&field)
	}

	c.Redirect(http.StatusFound, "/templates")
}

func GetTemplates(c *gin.Context) {
	email, _ := c.Get("email")
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	db.Where("email = ?", email).First(&user)

	var templates []models.Template
	db.Where("user_id = ?", user.ID).Find(&templates)

	c.HTML(http.StatusOK, "templates.html", gin.H{
		"Templates": templates,
	})
}

func GetTemplate(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var template models.Template
	db.Preload("Fields").First(&template, id)

	c.HTML(http.StatusOK, "template.html", gin.H{
		"Template": template,
	})
}

func UpdateTemplate(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var template models.Template
	db.First(&template, id)

	if err := c.ShouldBind(&template); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&template)

	c.Redirect(http.StatusFound, "/templates/"+id)
}

func DeleteTemplate(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var template models.Template
	db.First(&template, id)
	db.Delete(&template)

	c.Redirect(http.StatusFound, "/templates")
}

func NewTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "new_template.html", nil)
}

func EditTemplate(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var template models.Template
	db.Preload("Fields").First(&template, id)

	c.HTML(http.StatusOK, "edit_template.html", gin.H{
		"Template": template,
	})
}
