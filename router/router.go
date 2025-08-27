package router

import (
	"myblog-backend/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/classes", service.GetClasses)
	}

	return r
}
