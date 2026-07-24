package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
	IsLogin   bool
	LoginUser interface{}
}

func (this *BaseController) Prepare() {

	loginUser := this.GetSession("loginUser")

	fmt.Println("loginUser------->", loginUser)

	if loginUser != nil {
		this.IsLogin = true
		this.LoginUser = loginUser
	} else {
		this.IsLogin = false
	}

	this.Data["IsLogin"] = this.IsLogin

}
