package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"upv.life/server/common"
	"upv.life/server/config"
	"upv.life/server/db"
	"upv.life/server/middleware"
	"upv.life/server/model"
	"upv.life/server/service"
)

type LoginOrRegisterForm struct {
	Name string `json:"name" binding:"required"`
	Pwd  string `json:"pwd" binding:"required"`
}

type RegisterForm struct {
	Name string `json:"name" binding:"required,min=4,max=16"`
	Pwd  string `json:"pwd" binding:"required,min=6,max=20"`
}

func Register(c *gin.Context) {
	body := RegisterForm{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": common.Translate(err),
		})
		return
	}
	user := &model.User{
		Name:     body.Name,
		Nickname: body.Name,
		Pwd:      body.Pwd,
	}
	if user, token, error := service.Register(user); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": error.Error(),
		})
	} else {
		c.SetCookie("access_token", token, int(time.Now().Add(time.Hour*24*7).Second()), "/", config.Domain, false, false)
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
			"err": common.Translate(err),
		})
		return
	}

	user, err := service.GetUserByName(body.Name)
	if err == gorm.ErrRecordNotFound {
		if user, token, error := service.Register(&model.User{
			Name:     body.Name,
			Nickname: body.Name,
			Pwd:      body.Pwd,
			Email:    "",
		}); error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": error.Error(),
			})
		} else {
			c.SetCookie("access_token", token, 3600*24*7, "/", config.Domain, false, false)
			c.JSON(http.StatusOK, gin.H{
				"data": gin.H{
					"user":         user,
					"access_token": token,
				},
			})
		}
		return
	}

	if !common.ComparePasswords(user.Pwd, []byte(body.Pwd)) {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Invalid password.",
		})
		return
	}

	user.Pwd = ""
	token, err := middleware.GenerateJwtToken(user.ID, user.Name, user.Level)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.SetCookie("access_token", token, 3600*24*7, "/", config.Domain, false, false)
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
	if err := db.Orm.Where("id = ?", uid).First(&user).Error; err != nil {
		code := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			code = http.StatusUnauthorized
		}
		c.JSON(code, gin.H{
			"err": err.Error(),
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
	user, err := service.GetUserByName(name)
	if err == gorm.ErrRecordNotFound {
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
			"err": common.Translate(err),
		})
		return
	}

	var users []model.User
	tx := db.Orm.Scopes(model.Paginate(c))

	if body.Level != nil && *body.Level != 0 {
		tx.Where("level = ?", body.Level)
	}

	if body.Status != nil {
		tx.Where("status = ?", body.Status)
	}

	if body.Keyword != "" {
		tx.Where("name LIKE ? OR nickname LIKE ?", "%"+body.Keyword+"%", "%"+body.Keyword+"%")
	}

	tx.Order("created_at DESC").Find(&users)

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
	user, err := service.GetUserByName(userName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	if !common.IsRoot(ctxUser.Level) && !common.IsAdmin(ctxUser.Level) {
		if user.Status != body.Status || user.Level != body.Level || userName != ctxUser.Name {
			c.JSON(http.StatusForbidden, gin.H{
				"err": "Forbidden.",
			})
		}
	} else {
		if
		// 低等级不能修改高等级只能修改自己
		(ctxUser.UserId != user.ID && body.Level <= ctxUser.Level) ||
			//除了ROOT 其他角色不能修改自己的等级和状态
			(!common.IsRoot(ctxUser.Level) && ctxUser.UserId == user.ID &&
				(body.Status != user.Level || body.Level != user.Level)) {

			c.JSON(http.StatusForbidden, gin.H{
				"err": "Forbidden.",
			})
			return
		}
	}

	body.Name = ""
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

func GetUserStat(c *gin.Context) {
	ctxUser, _ := c.Get(middleware.CTX_AUTH_KEY)
	uid := uint(ctxUser.(*middleware.AuthClaims).UserId)
	var userStat model.UserStat

	// TODO: 重用posts表 注释查出来是多条数据
	err := db.Orm.Raw(`SELECT
	(SELECT COUNT(l.id) FROM likes l LEFT JOIN posts on posts.uid = ? WHERE l.pid IN(posts.id)) as LikesCount,
	(SELECT COUNT(c.id) FROM collections c LEFT JOIN posts on posts.uid = ? WHERE c.pid IN(posts.id)) as CollectionCount,
	(SELECT COUNT(cm.id) FROM comments cm LEFT JOIN posts on posts.uid = ? WHERE cm.pid IN(posts.id)) as CommentCount,
	(SELECT SUM(pr.hits) FROM post_rankings pr LEFT JOIN posts on posts.uid = ? AND posts.type = "post" WHERE pr.pid IN(posts.id)) as Pits,
	(SELECT SUM(vr.hits) FROM post_rankings vr LEFT JOIN posts on posts.uid = ? AND posts.type = "video" WHERE vr.pid IN(posts.id)) as Vits
	`, uid, uid, uid, uid, uid).Scan(&userStat).Error

	// err = db.Orm.Model(&model.Post{}).
	// 	Select(`
	// 	posts.id,
	// 	(SELECT COUNT(id) FROM likes WHERE likes.pid = posts.id) as LikesCount,
	// 	(SELECT COUNT(id) FROM collections WHERE collections.pid = posts.id) as CollectionCount,
	// 	(SELECT COUNT(id) FROM comments WHERE comments.pid = posts.id) as CommentCount,
	// 	(SELECT SUM(hits) FROM post_rankings WHERE pid = posts.id) as Hits
	// 	`).
	// 	Where("posts.uid = ?", uid).
	// 	Group("posts.type, posts.id").
	// 	Scan(&userStat).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userStat,
	})
}

func GetUserPostActivity(c *gin.Context) {
	ctxUser, _ := c.Get(middleware.CTX_AUTH_KEY)
	uid := uint(ctxUser.(*middleware.AuthClaims).UserId)
	var activities []model.PostActicity
	if err := db.Orm.Raw(`
	SELECT 'like' as type,posts.type as PostType,'' as Comment, posts.id as pid, posts.title as PostTitle, likes.created_at as CreatedAt, users.name as UserName,users.nickname as UserNickname,users.avatar as UserAvatar from likes LEFT JOIN posts on posts.id = likes.pid LEFT JOIN users on users.id = posts.uid WHERE posts.uid = ?
	UNION ALL
	SELECT 'collection',posts.type,'',posts.id, posts.title, collections.created_at, users.name,users.nickname,users.avatar from collections LEFT JOIN posts on posts.id = collections.pid LEFT JOIN users on users.id = posts.uid WHERE posts.uid = ?
	UNION ALL
	SELECT 'comment',posts.type, comments.content, posts.id, posts.title, comments.created_at, users.name,users.nickname,users.avatar from comments LEFT JOIN posts on posts.id = comments.pid LEFT JOIN users on users.id = posts.uid WHERE posts.uid = ?
	ORDER BY CreatedAt DESC LIMIT 15`, uid, uid, uid).Scan(&activities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": activities,
	})
}
