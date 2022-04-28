package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"upv.life/server/common"
	"upv.life/server/config"
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
	if user, error := service.Register(user); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": error.Error(),
		})
	} else {
		token, err := middleware.GenerateJwtToken(user.ID, user.Name, user.Level)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.SetCookie("access_token", token, int(time.Now().Add(time.Hour*-1).Second()), "/", config.Domain, false, false)
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"user":         user,
				"access_token": token,
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
	token, err := middleware.GenerateJwtToken(user.ID, user.Name, user.Level)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.SetCookie("access_token", token, int(time.Now().Add(time.Hour*-1).Second()), "/", config.Domain, false, false)
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user":         user,
			"access_token": token,
		},
	})
}

func GetUser(c *gin.Context) {
	ctxUser, _ := c.Get(middleware.CTX_AUTH_KEY)
	uid := uint(ctxUser.(*middleware.AuthClaims).UserId)
	var user model.User
	if err := db.Orm.Where("id = ?", uid).Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user": user,
		},
	})
}

func GetUserByName(c *gin.Context) {
	name := c.Param("name")
	var user model.User
	db.Orm.Where("name = ?", name).Find(&user)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "User not found.",
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
	tx := db.Orm.Scopes(model.Paginate(c))

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

func UpdateUserByName(c *gin.Context) {
	body := &model.User{}
	userName := c.Param("name")
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": common.Translate(err),
		})
		return
	}

	ctxUser := c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims)
	if !common.IsRoot(ctxUser.Level) && (userName != ctxUser.Name || body.Status != 0 || body.Level != 0) {
		c.JSON(http.StatusForbidden, gin.H{
			"err": "Forbidden.",
		})
		return
	}

	if !service.IsUserExistByName(userName) {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "User not found.",
		})
		return
	}

	if body.Name != "" {
		if _, err := service.CheckUserName(body.Name); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
	}

	if body.Pwd != "" && common.CheckPassword(body.Pwd) {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Invalid password.",
		})
	}

	if err := db.Orm.Model(body).Where("name = ?", c.Param("name")).Updates(body).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"err": nil})
}
