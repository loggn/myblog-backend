package service

import (
	"myblog-backend/config"
	"myblog-backend/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetClasses(c *gin.Context) {
	var classes []model.Class
	config.DB.Find(&classes)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": classes})
}

func GetArticlesByClassID(c *gin.Context) {
	classIDStr := c.Param("class_id")      //这里的Param（大概知道，但是顺便复习一下）
	classID, _ := strconv.Atoi(classIDStr) //这里为什么要转成int类型？而且，为什么传过来的时候不是int类型？

	var articles []model.Article
	config.DB.Where("class_id = ?", classID).Find(&articles)

	res := make([]map[string]interface{}, 0) //哦哦，这里的map是指的键值对是吗？
	for _, a := range articles {
		res = append(res, map[string]interface{}{
			"id":       a.ID,
			"class_id": a.ClassID,
			"title":    a.Title,
			"date":     a.Date,
			"tags":     strings.Split(a.Tags, ","), //以,号分割
			"file":     a.File,
		})
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": res})
}

func GetArticleByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var article model.Article
	config.DB.Where("id = ?", id).Find(&article)

	c.JSON(200, gin.H{
		"data": article,
	})
}

func CreateClass(c *gin.Context) {
	var req model.Class
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "无效请求"})
		return
	}

	if err := config.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "添加类别失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": req})
}

func CreateArticle(c *gin.Context) {
	var req model.Article
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "无效请求"})
		return
	}

	if err := config.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "添加文章失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": req})
}
