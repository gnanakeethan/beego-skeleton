package app

import (
	"backend/app/auth"
	_ "backend/app/auth/router"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/otiai10/copy"
)

func InitBeego() {
	_ = beego.AddFuncMap("hasPermission", auth.HasPermission)
	_ = beego.AddFuncMap("hasRole", auth.HasRole)
	err := copy.Copy("./app/views/", "./"+beego.BConfig.WebConfig.ViewsPath+"/honeypot/")
	if err != nil {
		panic(err)
	}
	beego.BConfig.EnableGzip = true
	beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js", ".html"}
}
