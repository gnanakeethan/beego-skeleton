package router

import (
	"backend/app/auth/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	authNamespace := beego.NewNamespace("/auth",
		beego.NSRouter("/login", &controllers.AuthController{}, "Get:Login"),
		beego.NSRouter("/login", &controllers.AuthController{}, "Post:LoginPost"),
		beego.NSRouter("/logout", &controllers.AuthController{}, "*:Logout"),
	)
	beego.AddNamespace(authNamespace)
}
