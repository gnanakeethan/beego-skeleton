package router

import (
	"backend/app/auth/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/auth/login", &controllers.AuthController{}, "Get:Login")
	beego.Router("/auth/login", &controllers.AuthController{}, "Post:LoginPost")
	beego.Router("/auth/logout", &controllers.AuthController{}, "*:Logout")
}
