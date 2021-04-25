package routers

import (
	"backend/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	// beego.Router("/auth/login", &controllers2.AuthController{},"Get:Login")
	// beego.Router("/auth/login", &controllers2.AuthController{},"Post:LoginPost")
	// beego.Router("/auth/logout", &controllers2.AuthController{},"*:Logout")
	// beego.AddNamespace(loginNamespace)
	// appNamespace := beego.NewNamespace("/app",
	// 	beego.NSInclude(
	// 		&controllers.AppController{},
	// 	),
	// )
	// beego.AddNamespace(appNamespace)
}
