package config
//
//import (
//	"os"
//	"fmt"
//	"io/ioutil"
//	"strings"
//	"github.com/json-iterator/go"
//	"gopkg.in/yaml.v2"
//)
//
//var registeredConfig = make(map[string]interface{})
//var AlreadyInitConfig = false
//
//func InitConfig(fn string) {
//	fd, err := os.Open(fn)
//	if err != nil {
//		fmt.Println("open config file error:", err.Error())
//	}
//	defer fd.Close()
//
//	content, err := ioutil.ReadAll(fd)
//	if err != nil {
//		fmt.Println("read config file error:", err.Error())
//	}
//
//	var m map[string]interface{}
//	var tagName string
//	if strings.HasSuffix(fn, "json") {
//		tagName = "json"
//		if err = jsoniter.Unmarshal(content, &m); err != nil {
//			fmt.Println("Unmarshal config error:")
//		}
//	} else if strings.HasSuffix(fn, "yaml") {
//		tagName = "yaml"
//		if err = yaml.Unmarshal(content, &m); err != nil {
//			fmt.Println("Unmarshal config error:")
//		}
//	}
//	for k,conf := range registeredConfig{
//		if value,ok :=m[k]; ok{
//			decoder,_ :=mapst
//		}
//	}
//}
