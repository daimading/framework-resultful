package handlers

import "github.com/gin-gonic/gin"

func InitUserRouter(router gin.IRouter) {
	userGroup := router.Group("/user")
	{
		//注册路由
		userGroup.Any("/ping", Ping)
		userGroup.POST("/register", Register)
	}
}
