package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"upv.life/server/config"
)

var (
	SMMSToken          = ""
	RetryCount         = 0
	ImageMaxSize int64 = 5 << 20
	FileMaxSize  int64 = 700 << 20
	UploadPath         = "./uploads/"
)

func SMMSImageUploder(c *gin.Context) {
	c.Request.ParseMultipartForm(ImageMaxSize)
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "parse file error",
		})
		return
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".gif" && ext != ".jpg" && ext != ".png" && ext != ".jpeg" && ext != ".webp" && ext != ".bmp" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "file type not allowed",
		})
		return
	}

	token := _SMMSAuth()
	if token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "uploader token error",
		})
		return
	}

	if url, err := _SMMSUploder(file, header, token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": url,
		})
	}
}

func FileUploader(c *gin.Context) {
	c.Request.ParseMultipartForm(FileMaxSize)
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "parse file error",
		})
		return
	}
	defer file.Close()

	prefix := UploadPath + time.Now().Format("2006/01/02") + "/"
	if _, err := os.Stat(prefix); os.IsNotExist(err) {
		// mask := syscall.Umask(0)
		if err := os.MkdirAll(prefix, 0777); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
			// defer syscall.Umask(mask)
			return
		}
		// defer syscall.Umask(mask)
	}

	filename := header.Filename
	ext := strings.ToLower(filepath.Ext(filename))
	filename = strconv.FormatInt(time.Now().UnixNano(), 10) + ext

	f, err := os.OpenFile(prefix+filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	defer f.Close()

	if _, err := io.Copy(f, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": prefix + filename,
		})
	}
}

func TencentCOSUploader(c *gin.Context) {
}

/////////////////////////////////////////////////////////////////////////////////////////////

func _SMMSUploder(file multipart.File, fileHeader *multipart.FileHeader, token string) (string, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	bodyWriter.WriteField("format", "json")
	fileWriter, _ := bodyWriter.CreateFormFile("smfile", fileHeader.Filename)
	io.Copy(fileWriter, file)

	req, _ := http.NewRequest("POST", "https://sm.ms/api/v2/upload", bodyBuf)
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", bodyWriter.FormDataContentType())

	response, _ := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(response.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	defer response.Body.Close()
	defer bodyWriter.Close()

	if data["success"].(bool) {
		return data["data"].(map[string]interface{})["url"].(string), nil
	} else {
		// retry once
		if data["code"] == "unauthorized" && RetryCount == 0 {
			fmt.Print("unauthorized, retry", RetryCount)
			RetryCount = 1
			_SMMSAuth()
			s, e := _SMMSUploder(file, fileHeader, token)
			RetryCount = 0
			return s, e
		}
		return "", errors.New(data["message"].(string))
	}
}

func _SMMSAuth() string {
	if SMMSToken != "" {
		return SMMSToken
	}

	if response, err := http.Post(`https://sm.ms/api/v2/token?username=`+config.SMMSUserName+`&password=`+config.SMMSPassword, "application/json", nil); err != nil {
		return ""
	} else {
		body, _ := io.ReadAll(response.Body)
		var data map[string]interface{}
		json.Unmarshal(body, &data)
		return data["data"].(map[string]interface{})["token"].(string)
	}
}

// build linux,386 darwin,!cgo
