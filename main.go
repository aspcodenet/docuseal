package main

import (
	"github.com/gin-gonic/gin"
	"github.com/user/repo/handlers"
	"github.com/user/repo/middleware"
	"github.com/user/repo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{}, &models.Template{}, &models.Submission{}, &models.Field{}, &models.AccessToken{})

	r := gin.Default()

	// Add the database to the context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", handlers.Ping)
	r.POST("/users", handlers.CreateUser)
	r.POST("/login", handlers.Login)

	authRequired := r.Group("/")
	authRequired.Use(middleware.AuthMiddleware())
	{
		authRequired.GET("/dashboard", handlers.Dashboard)
		authRequired.GET("/templates", handlers.GetTemplates)
		authRequired.GET("/templates/new", handlers.NewTemplate)
		authRequired.POST("/templates", handlers.CreateTemplate)
		authRequired.GET("/templates/:id", handlers.GetTemplate)
		authRequired.GET("/templates/:id/edit", handlers.EditTemplate)
		authRequired.POST("/templates/:id", handlers.UpdateTemplate)
		authRequired.POST("/templates/:id/delete", handlers.DeleteTemplate)

		authRequired.GET("/templates/:id/submissions/new", handlers.NewSubmission)
		authRequired.POST("/templates/:id/submissions", handlers.CreateSubmission)
		authRequired.GET("/submissions/:id", handlers.GetSubmission)
	}

	api := r.Group("/api")
	api.Use(middleware.APIAuthMiddleware())
	{
		api.GET("/templates", handlers.APIGetTemplates)
		api.GET("/templates/:id", handlers.APIGetTemplate)
		api.POST("/templates", handlers.APICreateTemplate)
		api.PUT("/templates/:id", handlers.APIUpdateTemplate)
		api.DELETE("/templates/:id", handlers.APIDeleteTemplate)

		api.GET("/submissions", handlers.APIGetSubmissions)
		api.GET("/submissions/:id", handlers.APIGetSubmission)
		api.POST("/submissions", handlers.APICreateSubmission)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
