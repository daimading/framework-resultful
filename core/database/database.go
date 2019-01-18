package database

import (
	"github.com/jinzhu/gorm"
	"sync"
	"fmt"
)

var db *gorm.DB
var once sync.Once

//获取数据库连接
func GetDB() *gorm.DB {
	once.Do(loadDb)
	return db
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
