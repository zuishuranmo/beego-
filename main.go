package main

import (
	"DataCertProject/db_mysql"
	_ "DataCertProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	//连接数据库
	db_mysql.ConnectDB()
	//静态资源路径设置

	beego.Run()
}

