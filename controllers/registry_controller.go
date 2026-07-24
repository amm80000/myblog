package controllers

import (
	"fmt"
	"myblog/models"
	"myblog/utils"
	"time"
)

type RegistryController struct {
	BaseController
}

func (c *RegistryController) Get() {
	c.TplName = "register.html"
}

// 处理注册

func (c *RegistryController) Post() {
	// 获取表单信息
	username := c.GetString("username")
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	fmt.Println(username, password, repassword)

	// 注册之前先判断该用户是否已经被注册，如果已经注册，返回错误
	id := models.QueryUserWithUsername(username)
	fmt.Println("id:", id)

	if id > 0 {

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名已经存在"}
		c.ServeJSON()
		return
	}

	// 注册用户名和密码
	var md5Password = utils.MD5(password)
	fmt.Println("md5后：", md5Password)

	user := models.Users{0, username, md5Password, 0, time.Now().Unix()}

	_, err := models.InsertUser(user)

	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	}

	c.ServeJSON()
}
