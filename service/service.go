package service

import (
	"myblog-backend/data"

	"github.com/gin-gonic/gin"
)

func GetClasses(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": data.Classes,
	})
}
