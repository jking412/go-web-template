package main

import (
	"github.com/gin-gonic/gin"
	"go-web-template/api"
	"go-web-template/boot"
	"go-web-template/conf"
	"go-web-template/pkg/logger"
)

func main() {
	boot.Initialize()

	r := gin.Default()

	api.RegisterApi(r)

	if err := r.Run(":" + conf.ServerPort); err != nil {
		logger.ErrorString("web", "启动失败", err.Error())
	}
}
