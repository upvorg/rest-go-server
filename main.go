package main

import (
	"github.com/gin-gonic/gin"
	"upv.life/server/common"
	"upv.life/server/config"
	"upv.life/server/controller"
	"upv.life/server/db"
	"upv.life/server/middleware"
)

func main() {
	config.Initialize()
	gin.SetMode(config.AppMode)

	r := gin.Default()
	r.SetTrustedProxies([]string{"::"})

	// initailize database.
	db.Initialize()

	// setup auth middleware.
	r.Use(middleware.Auth())

	common.InitValidator()

	// setup routes.
	controller.Initialize(r)

	r.Run(":" + config.AppPort)
}
