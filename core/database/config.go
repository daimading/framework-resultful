package database

import (
	"bytes"
	"fmt"
	"strconv"
)

type DBConfig struct {
	DriverName      string `json:"driver_name" yaml:"driver_name"`
	DBName          string `json:"db_name" yaml:"driver_name"`
	Host            string `json:"host" yaml:"driver_name"`
	Port            int    `json:"port" yaml:"driver_name"`
	UserName        string `json:"user_name" yaml:"driver_name"`
	Password        string `json:"password" yaml:"driver_name"`
	MaxOpenConn     int    `json:"max_open_conn" yaml:"driver_name"`
	MaxIdleConn     int    `json:"max_idle_conn" yaml:"driver_name"`
	ConnMaxLifeTime int    `json:"conn_max_life_time" yaml:"driver_name"`
}

func (dbConfig *DBConfig) String() string {
	var conStr bytes.Buffer
	switch dbConfig.DriverName {
	case "sqlite", "sqlite3":
		return dbConfig.DBName
	case "postgres", "postgresql":
		dbName := "host=%s post=%d user=%s password=%s dbname=%s sslmode=%s"
		return fmt.Sprintln(dbName, dbConfig.Host, dbConfig.Port, dbConfig.UserName, dbConfig.Password, dbConfig.DBName)
	case "mysql":
		//<user>:<password>/<database>?charset=utf8&parseTime=True&loc=Local
		conStr.WriteString(dbConfig.UserName)
		conStr.WriteString(":")
		conStr.WriteString(dbConfig.Password)
		conStr.WriteString("@tcp(")
		conStr.WriteString(dbConfig.Host)
		conStr.WriteString(":")
		conStr.WriteString(strconv.Itoa(dbConfig.Port))
		conStr.WriteString(")/")
		conStr.WriteString(dbConfig.DBName)
		conStr.WriteString("?charset=utf8&parseTime=true")
	}
	return conStr.String()
}

var Config = DBConfig{
	DriverName: "sqlite3",
	DBName:     "test.sqlite3",
}

//TODO 配置数据库表注册
func init() {
}
