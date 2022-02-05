package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Masterminds/squirrel"
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
		token, err := middleware.GenerateJwtToken(user.ID, user.Name)
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
	token, err := middleware.GenerateJwtToken(user.ID, user.Name)
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
	Level   string `form:"level,omitempty"`
}

type PaginationUser struct {
	model.Pagination
	Data []*model.User `json:"data"`
}

func GetUsers(c *gin.Context) {
	body := GetUsersForm{}
	if err := c.ShouldBindQuery(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": common.Translate(err),
		})
		return
	}

	sq := squirrel.Select("*").From("users")
	if body.Keyword != "" {
		like := "%" + body.Keyword + "%"
		sq = sq.Where(squirrel.Or{
			squirrel.Like{"name": like},
			squirrel.Like{"nickname": like},
		})
	}
	if body.Level != "" {
		sq = sq.Where(squirrel.Eq{"level": body.Level})
	}

	query, arg, _ := sq.Offset(body.Limit * (body.Page - 1)).Limit(body.Limit).ToSql()
	countQuery, arg, _ := sq.ToSql()

	var count uint64
	db.Sqlx.Select(count, countQuery, arg...)
	users := []*model.User{}
	if err := db.Sqlx.Select(&users, query, arg...); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": PaginationUser{
			Pagination: model.Pagination{
				Page:  body.Page,
				Limit: body.Limit,
				Total: count/body.Limit + 1,
			},
			Data: users,
		}})

}
