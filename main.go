package main

import (
	_ "BeegoDemo2/routers"
	"BeegoDemo2/utils"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	utils.InitMysql()
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
