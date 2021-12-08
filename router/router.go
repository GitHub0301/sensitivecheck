package router

import (
	"github.com/astaxie/beego"
	"sensitivecheck/controllers"
)

func init() {
	beego.Router("/", &controllers.SensitiveCheckController{})
	beego.Router("/check", &controllers.SensitiveCheckController{})
}
