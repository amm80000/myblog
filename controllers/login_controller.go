package controllers

import (
	"fmt"
	"myblog/models"
	"myblog/utils"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

// 登录
func (this *LoginController) Post() {

	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Printf("username:%s,password:%s\n", username, password)
	id := models.QueryUserParam(username, utils.MD5(password))
	fmt.Printf("id:%d\n", id)
	if id > 0 {
		this.SetSession("loginUser", username)
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}

	this.ServeJSON()

}
