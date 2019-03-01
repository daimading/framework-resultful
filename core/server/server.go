package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var router gin.IRouter
var once sync.Once
var registerRouteFuncs []func(gin.IRouter)

func Start() error {
	engine := GetRouter().(*gin.Engine)

	//此时注册所有路由
	for _, registerRoute := range registerRouteFuncs {
		registerRoute(engine)
	}
	//server.InitDataBase(appConfig)
	if len(Config.CertPath) != 0 && len(Config.KeyPath) != 0 {
		return engine.RunTLS(fmt.Sprintf("%s:%d", Config.IP, Config.Port),
			Config.CertPath, Config.KeyPath)
	} else {
		return engine.Run(fmt.Sprintf("%s:%d", Config.IP, Config.Port))
	}
}

func GetRouter() gin.IRouter {
	once.Do(initRouter)
	return router
}
func initRouter() {
	router = gin.Default()
}
func RegisterRoute(iRouter ...func(gin.IRouter)) {
	registerRouteFuncs = append(registerRouteFuncs, iRouter...)
}
