package app

import (
	"backend/app/auth"
	_ "backend/app/auth/router"
	"backend/app/lang"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
	"github.com/kr/pretty"
	"github.com/otiai10/copy"
	"gopkg.in/hlandau/passlib.v1"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func InitBeego() {
	sslmode := "?sslmode=" + beego.AppConfig.DefaultString("ssl_mode", "disable")

	dbdriver := beego.AppConfig.DefaultString("dbdriver", "mysql")
	dbdriverPrefix := beego.AppConfig.DefaultString("dbdriver_prefix", "")
	conn := dbdriverPrefix + beego.AppConfig.DefaultString(dbdriver+"_user", "root") + ":" +
		beego.AppConfig.DefaultString(dbdriver+"_pass", "") + "@" +
		beego.AppConfig.DefaultString(dbdriver+"_urls", "tcp(127.0.0.1:3306)") + "/" +
		beego.AppConfig.DefaultString(dbdriver+"_db", "microlearn") + sslmode
	pretty.Println(conn)
	_ = orm.RegisterDataBase("default", dbdriver, conn)
	_ = passlib.UseDefaults(passlib.DefaultsLatest)

	// template function for checking permission, can be cached
	_ = beego.AddFuncMap("hasPermission", auth.HasPermission)
	_ = beego.AddFuncMap("hasRole", auth.HasRole)

	// load language
	lang.LoadLang()

	// set translation function
	_ = beego.AddFuncMap("_", i18n.Tr)

	// copy view files
	err := copy.Copy("./app/views/", "./"+beego.BConfig.WebConfig.ViewsPath+"/honeypot/")
	if err != nil {
		panic(err)
	}
	beego.BConfig.EnableGzip = true
	beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js", ".html"}
	orm.Debug = true
}
