package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"upv.life/server/common"
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
	user := &model.User{
		Name: body.Name,
		Pwd:  body.Pwd,
	}
	if error := user.GetUserByNameWithPwd(); error != nil {
		if error == sql.ErrNoRows {
			Register(c)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": error.Error(),
			})
		}
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

func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := &model.User{
		ID: uint64(id),
	}
	if err := user.GetUserByID(); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "User not found.",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

type GetUsersForm struct {
	model.Pagination
	Keyword string `form:"keyword,omitempty"`
	Level   *uint8 `form:"level,omitempty"`
	Status  *uint8 `form:"status,omitempty"`
}

func GetUsers(c *gin.Context) {
	body := GetUsersForm{}
	if err := c.ShouldBindQuery(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": common.Translate(err),
		})
		return
	}

	user := &model.User{Name: body.Keyword, Level: body.Level, Status: body.Status}
	users, count, err := user.GetAll((body.Pagination.Page), (body.Pagination.Limit))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": model.PaginationData{
				Pagination: model.Pagination{
					Page:  body.Page,
					Limit: body.Limit,
					Total: count,
				},
				Data: users,
			}})
	}
}
