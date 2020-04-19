package conf

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/joho/godotenv"
	"online_shop/model"
	"os"
)

func Init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(".env load error")
	}
	initLogs()
	fmt.Println(os.Getenv("MYSQL_DSN"))
	model.DateBase(os.Getenv("MYSQL_DSN"))
}

func initLogs() {
	logs.SetLogger("file", `{"filename":"logs/log.log"}`)
	logs.EnableFuncCallDepth(true)

}
