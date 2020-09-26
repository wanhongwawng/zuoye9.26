package routers

import (
	"beego02/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//register:注册
    beego.Router("/register",&controllers.RegisterController{})
    //http://127.0.0..8080
    beego.Router("/",&controllers.MainController{})
}
