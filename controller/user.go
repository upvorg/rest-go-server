package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"upv.life/server/common"
	"upv.life/server/db"
	"upv.life/server/middleware"
	"upv.life/server/model"
	"upv.life/server/service"
)

type LoginOrRegisterForm struct {
	Name string `json:"name" binding:"required,min=4,max=16"`
	Pwd  string `json:"pwd" binding:"required,min=6,max=20"`
}

func Register(c *gin.Context) {
	body := LoginOrRegisterForm{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": common.Translate(err),
		})
		return
	}
	user := &model.User{
		Name:     body.Name,
		Nickname: body.Name,
		Pwd:      body.Pwd,
	}
	if _, error := service.Register(user); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": error.Error(),
		})
	} else {
		token, err := middleware.GenerateJwtToken(uint64(user.ID), user.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"token": token,
				"user":  user,
			},
		})
	}
}

func Login(c *gin.Context) {
	body := LoginOrRegisterForm{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": common.Translate(err),
		})
		return
	}

	var user model.User
	db.Orm.Where(&model.User{Name: body.Name}).Find(&user)

	if user.ID == 0 {
		Register(c)
		return
	}

	if !common.ComparePasswords(user.Pwd, []byte(body.Pwd)) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid password.",
		})
		return
	}
	user.Pwd = ""
	token, err := middleware.GenerateJwtToken(uint64(user.ID), user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"token": token,
			"user":  user,
		},
	})
}

func GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user model.User
	db.Orm.Where("id = ?", id).Find(&user)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "User not found.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

type GetUsersForm struct {
	Keyword string `form:"keyword,omitempty"`
	Level   *int   `form:"level,omitempty"`
	Status  *int   `form:"status,omitempty"`
}

func GetUsers(c *gin.Context) {
	body := GetUsersForm{}
	if err := c.ShouldBindQuery(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": common.Translate(err),
		})
		return
	}

	var users []model.User
	tx := db.Orm.Debug().Scopes(model.Paginate(c))

	if body.Level != nil {
		tx.Where("level = ?", body.Level)
	}

	if body.Status != nil {
		tx.Where("status = ?", body.Status)
	}

	if body.Keyword != "" {
		tx.Where("name LIKE ? OR nickname LIKE ?", "%"+body.Keyword+"%", "%"+body.Keyword+"%")
	}

	tx.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
