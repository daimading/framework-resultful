package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var db *gorm.DB
var once sync.Once
var tables = make([]interface{}, 0)

//获取数据库连接
func GetDB() *gorm.DB {
	once.Do(loadDb)
	return db
}

//初始化数据库，创建连接
func Init() {
	logrus.Info("database init")
}

//初始化数据库，创建连接
//TODO 初始化表，注册表
func loadDb() {
	var err error
	dbConfig := Config

	if db != nil {
		db.Close()
		db = nil
	}
	db, err = gorm.Open(dbConfig.DriverName, dbConfig.String())
	if err != nil {
		fmt.Println("database error:")
	}

	if err = db.DB().Ping(); err != nil {
		CloseDB()
		fmt.Println("ping error: ", err.Error())
	}

	db.SingularTable(true)

	if dbConfig.DriverName != "sqlite3" {
		if dbConfig.MaxIdleConn != 0 {
			db.DB().SetMaxIdleConns(dbConfig.MaxIdleConn)
		}
		if dbConfig.MaxIdleConn != 0 {
			db.DB().SetMaxOpenConns(dbConfig.MaxIdleConn)
		}
		if dbConfig.ConnMaxLifeTime != 0 {
			db.DB().SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifeTime) * time.Second)
		}
	}
}

//关闭数据库连接
func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			fmt.Println("database connection closed!!!")
		}
	}
}

//注册表
func RegisterTable(value interface{}) {
	fmt.Println("register table:", value)
}
func MigrateAllTable() error {
	if len(tables) > 0 {
		return db.AutoMigrate(tables...).Error
	}
	return nil
}
