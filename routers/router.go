package routers

import (
	"backend/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	appNamespace := beego.NewNamespace("/app",
		beego.NSInclude(&controllers.AppController{}),
	)
	beego.AddNamespace(appNamespace)
}
