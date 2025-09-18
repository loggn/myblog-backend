package router

import (
	"myblog-backend/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	api := r.Group("/api")
	{
		api.GET("/classes", service.GetClasses)
		api.GET("/articles/class/:class_id", service.GetArticlesByClassID)
		api.GET("/article/:id", service.GetArticleByID)

		api.POST("/class", service.CreateClass)
		api.POST("/article", service.CreateArticle)
	}

	return r
}
