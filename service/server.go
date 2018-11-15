package service

import (
	"github.com/astaxie/beego" // 使用beego框架
)

type MainController struct {
    beego.Controller  // 这里可以看做是其他语言中的继承
}

func (this *MainController) Get() {
	// 获取get请求的参数
	username := this.Ctx.Input.Param(":username") 
	// 输出字符串
	this.Ctx.WriteString("Hello " + username + "!")
}



