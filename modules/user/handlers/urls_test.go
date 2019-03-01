package handlers

import (
	"github.com/gin-gonic/gin"
	"testProject/core"
	"testProject/core/config"
)

var router *gin.Engine
var cookies string

func InitRouterConfig() {
	router = gin.Default()
	InitUserRouter(router)
}
func init() {
	config.InitConfig("../../../app.yaml")
	core.LoadCore()
	InitRouterConfig()
}
