package main

import (
	"github.com/gin-gonic/gin"
	"upv.life/server/config"
	"upv.life/server/controller"
	"upv.life/server/db"
	"upv.life/server/middleware"
)

func main() {
	config.Initialize()

	r := gin.Default()
	r.SetTrustedProxies([]string{"::"})

	// setup auth middleware.
	r.Use(middleware.Auth())

	// setup routes.
	controller.Initialize(r)

	// initailize database.
	db.Initialize()

	r.Run(":" + config.AppPort)
}
