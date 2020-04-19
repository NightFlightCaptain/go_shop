package model

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func DateBase(connectDB string) {
	db, err := gorm.Open("mysql", connectDB)
	if err != nil {
		logs.Error("数据库连接失败", err.Error())
	} else {
		logs.Info("连接数据库成功")
	}
	maxIdleConn, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
	if err != nil {
		logs.Error("DB_MAX_IDLE_CONN设置出错")
	}
	maxOpenConn, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
	if err != nil {
		logs.Error("DB_MAX_OPEN_CONN设置出错")
	}
	db.DB().SetMaxIdleConns(maxIdleConn)
	db.DB().SetMaxOpenConns(maxOpenConn)
	db.DB().SetConnMaxLifetime(30 * time.Second)

	db.SingularTable(true)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return os.Getenv("TABLE_PREFIX") + defaultTableName
	}
	db.AutoMigrate(&User{})
	DB = db
}
