package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"upv.life/server/common"
	"upv.life/server/model"
	"upv.life/server/service"
)

func Register(c *gin.Context) {
	body := struct {
		Name string `json:"name" binding:"required,min=4,max=16"`
		Pwd  string `json:"pwd" binding:"required,min=6,max=20"`
	}{}
	if error := c.ShouldBindJSON(&body); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": common.Translate(error),
		})
		return
	}
	user := &model.User{
		Name:     body.Name,
		Nickname: body.Name,
		Pwd:      body.Pwd,
	}
	if result, error := service.Register(user); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": error.Error(),
		})
	} else {
		user.ID = result.ID
		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	}
}
