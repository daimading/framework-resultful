package core

import (
	"net/http"
	"testProject/core/config"
	"testProject/core/database"
	"testProject/core/global"
)

var InitConfig = config.InitConfig

func LoadCore() {
	if config.AlreadyInitConfig {
		http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = global.Config.Keepalive
		global.Init()
		database.Init()
	}
}
