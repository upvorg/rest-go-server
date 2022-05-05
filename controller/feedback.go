package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"upv.life/server/common"
	"upv.life/server/db"
	"upv.life/server/middleware"
	"upv.life/server/model"
)

func CreateFeedback(c *gin.Context) {
	feedback := &model.Feedback{}
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": common.Translate(err),
		})
		return
	}

	if ctxUser, exists := c.Get(middleware.CTX_AUTH_KEY); exists {
		feedback.Name = ctxUser.(*middleware.AuthClaims).Name
	} else {
		feedback.Name = ""
	}

	feedback.Ip = c.ClientIP()
	err := db.Orm.Create(feedback).Error
	c.JSON(http.StatusOK, gin.H{
		"err": err,
	})
}

func GetFeedbacks(c *gin.Context) {
	feedbacks := []model.Feedback{}
	err := db.Orm.
		Order("created_at DESC").
		Find(&feedbacks).
		Error
	c.JSON(http.StatusOK, gin.H{
		"err":  err,
		"data": feedbacks,
	})
}
