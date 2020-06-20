package routers

import (
	"qrProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/form", &controllers.MainController{},"post:FormData")
	beego.Router("/tel", &controllers.MainController{},"get:TelData")
}
