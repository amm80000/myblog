package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	_ "myblog/routers"
	"myblog/utils"
)

func main() {

	utils.InitMysql()
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
