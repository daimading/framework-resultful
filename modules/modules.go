package modules

import (
	"testProject/core/server"
	user "testProject/modules/user/handlers"
)

func LoadModules() {
	//延迟注册route
	//避免在load modules时就初始化server实例
	server.RegisterRoute(
		user.InitUserRouter,
	)
}
