package controllers

import (
	"DataCertProject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post() {
	//1.解析请求数据
	var user models.User
	err :=r.ParseForm(&user)
	if err != nil{

		r.Ctx.WriteString("抱歉，数据解析错误，请重试！")
		return
	}
	//2.保存数据到数据库
	_, err = user.SaveUser()
	if err !=nil {
		fmt.Println(err)
		r.Ctx.WriteString("用户注册失败，请重试！")
		return
	}
	//用户注册成功
	r.TplName = "login.html"
}
