package main

import (
	"github.com/hath/cloudgo/service"
	flag "github.com/spf13/pflag"
	"github.com/astaxie/beego" // 使用beego框架
)

func main()  {
	// 命令行参数, 确定监听端口
	pPort := flag.StringP("port", "p", "8080", "the port to listen") 
	flag.Parse()
	// 路由设置, 把形如"/:username"的请求发到对应的控制器MainController执行相应逻辑
	beego.Router("/:username", &service.MainController{}) 
	// 启动应用, 监听对应的端口
	beego.Run("localhost:" + *pPort)
}

