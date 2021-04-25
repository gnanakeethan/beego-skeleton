package app

import (
	"backend/app/auth"
	_ "backend/app/auth/router"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/otiai10/copy"
	"gopkg.in/hlandau/passlib.v1"
)

func InitBeego() {
	sslmode := "require&sslrootcert=conf/ca.crt&sslkey=conf/client.root.key&sslcert=conf/client.root.crt"
	if beego.BConfig.RunMode == "dev" {
		sslmode = "disable"
	}
	dbdriver := beego.AppConfig.DefaultString("dbdriver", "")
	conn := "postgresql://" +
		beego.AppConfig.DefaultString(dbdriver+"_user", "") + ":" +
		beego.AppConfig.DefaultString(dbdriver+"_pass", "") + "@" +
		beego.AppConfig.DefaultString(dbdriver+"_urls", "") + "/" +
		beego.AppConfig.DefaultString(dbdriver+"_db", "") + "?sslmode=" + sslmode

	_ = orm.RegisterDataBase("default", dbdriver, conn)
	_ = passlib.UseDefaults(passlib.DefaultsLatest)
	_ = beego.AddFuncMap("hasPermission", auth.HasPermission)
	_ = beego.AddFuncMap("hasRole", auth.HasRole)
	err := copy.Copy("./app/views/", "./"+beego.BConfig.WebConfig.ViewsPath+"/honeypot/")
	if err != nil {
		panic(err)
	}
	beego.BConfig.EnableGzip = true
	beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js", ".html"}
}
