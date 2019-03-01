package config

import (
	"fmt"
	"github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

var registeredConfigs = make(map[string]interface{})
var AlreadyInitConfig = false

func InitConfig(fn string) {
	fd, err := os.Open(fn)
	if err != nil {
		fmt.Println("open config file error:", err.Error())
	}
	defer fd.Close()

	content, err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println("read config file error:", err.Error())
	}

	var m map[string]interface{}
	var tagName string
	if strings.HasSuffix(fn, "json") {
		tagName = "json"
		if err = jsoniter.Unmarshal(content, &m); err != nil {
			fmt.Println("Unmarshal config error:")
		}
	} else if strings.HasSuffix(fn, "yaml") {
		tagName = "yaml"
		if err = yaml.Unmarshal(content, &m); err != nil {
			fmt.Println("Unmarshal config error:")
		}
	}
	for k, conf := range registeredConfigs {
		if value, ok := m[k]; ok {
			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: tagName, Result: conf})
			if err = decoder.Decode(value); err != nil {
				fmt.Println("decoder config error:")
			}
		}
	}
	AlreadyInitConfig = true
}

func RegisterConfig(key string, config interface{}) {
	if _, ok := registeredConfigs[key]; ok {
	}
	if len(key) == 0 {
		fmt.Println("len", len(key))
	} else if config == nil {
		fmt.Println("config nil!!!!!")
	} else if _, ok := registeredConfigs[key]; ok {
		fmt.Println()
	}
}
