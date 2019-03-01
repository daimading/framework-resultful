package global

import (
	"testProject/core/config"
)

var Config = struct {
	LogLevel  string `json:"loglevel" yaml:"loglevel"`
	GCPercent int    `json:"gcpercent" yaml:"gcpercent"`
	Keepalive int    `json:"keepalive" yaml:"keepalive"`
	Debug     bool   `json:"debug" yaml:"debug"`
}{
	LogLevel:  "debug",
	GCPercent: 200,
	Keepalive: 30,
	Debug:     true,
}

func init() {
	config.RegisterConfig("global", &Config)
}
