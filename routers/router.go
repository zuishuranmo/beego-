package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)

/**
 * router.go文件的作以为：
 */

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user_register", &controllers.RegisterController{})
}
