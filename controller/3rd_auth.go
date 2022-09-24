package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"upv.life/server/config"
	"upv.life/server/db"
	"upv.life/server/middleware"
	"upv.life/server/model"
)

type QQAccessTokenResp struct {
	Access_token      string `json:"access_token"`
	Expires_in        int    `json:"expires_in"`
	Refresh_token     string `json:"refresh_token"`
	Error             int    `json:"error"`
	Error_description string `json:"error_description"`
}

type QQOpenIDResp struct {
	OpenID            string `json:"openid"`
	ClientID          string `json:"client_id"`
	Error             int    `json:"error"`
	Error_description string `json:"error_description"`
}

type QQGetUserInfoResp struct {
	Error             int    `json:"error"`
	Error_description string `json:"error_description"`
	Nickname          string `json:"nickname"`
}

func QQLogin(c *gin.Context) {
	if config.QQAppID == "" || config.QQAppKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "qq appid or appkey not set",
		})
		return
	}

	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "code is required",
		})
		return
	}

	// get access token
	url := "https://graph.qq.com/oauth2.0/token?grant_type=authorization_code&client_id=" + config.QQAppID +
		"&client_secret=" + config.QQAppKey +
		"&code=" + code +
		"&redirect_uri=http://" + config.Domain + "/qq_login" +
		"&fmt=json"
	resp, _ := http.Get(url)
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	var qqAccessTokenResp QQAccessTokenResp
	json.Unmarshal(body, &qqAccessTokenResp)
	if qqAccessTokenResp.Error != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": qqAccessTokenResp.Error,
			"err":  qqAccessTokenResp.Error_description,
		})
		return
	}

	// get openid
	url = "https://graph.qq.com/oauth2.0/me?access_token=" + qqAccessTokenResp.Access_token + "&fmt=json"
	resp, _ = http.Get(url)
	body, _ = io.ReadAll(resp.Body)
	defer resp.Body.Close()
	var qqOpenIDResp QQOpenIDResp
	json.Unmarshal(body, &qqOpenIDResp)
	if qqOpenIDResp.Error != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": qqOpenIDResp.Error,
			"err":  qqOpenIDResp.Error_description,
		})
		return
	}

	var retuernUser *model.User
	var retuenToken string
	err := db.Orm.Where("qq_openid = ?", qqOpenIDResp.OpenID).Find(retuernUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	if err == nil {
		//login
		token, err := middleware.GenerateJwtToken(retuernUser.ID, retuernUser.Name, retuernUser.Level)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
			return
		}
		retuenToken = token
	}

	if err == gorm.ErrRecordNotFound {
		// register
		url = "https://graph.qq.com/user/get_user_info?access_token=" + qqAccessTokenResp.Access_token +
			"&oauth_consumer_key=" + config.QQAppID + "&openid=" + qqOpenIDResp.OpenID
		resp, _ = http.Get(url)
		body, _ = io.ReadAll(resp.Body)
		defer resp.Body.Close()
		var qqGetUserInfoResp QQGetUserInfoResp
		json.Unmarshal(body, &qqGetUserInfoResp)
		if qqGetUserInfoResp.Error != 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": qqGetUserInfoResp.Error,
				"err":  qqGetUserInfoResp.Error_description,
			})
			return
		}
		u := map[string]interface {
		}{
			"Name":      qqGetUserInfoResp.Nickname,
			"Nickname":  qqGetUserInfoResp.Nickname,
			"qq_openid": qqOpenIDResp.OpenID,
			"Level":     4,
			"Status":    2,
		}

		if err := db.Orm.Model(&model.User{}).Create(&u).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
			return
		}
		token, _ := middleware.GenerateJwtToken(u["ID"].(uint), u["Name"].(string), u["Level"].(uint))
		retuenToken = token
	}

	c.SetCookie("access_token", retuenToken, 3600*24*7, "/", config.Domain, false, false)
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user":         retuernUser,
			"access_token": retuenToken,
		},
	})
}

func GoogleLogin(c *gin.Context) {}
