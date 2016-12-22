package routers

import (
	"LLB/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/query",&controllers.QueryController{})
	beego.Router("/newtask",&controllers.NewTaskController{})
	beego.Router("/starttask",&controllers.StartTaskController{})
	beego.Router("/addnews",&controllers.AddNewsController{})

}
