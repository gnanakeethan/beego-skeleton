package main

import (
	"fmt"

	"backend/app"
	_ "backend/routers"

	_ "github.com/lib/pq"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"gopkg.in/hlandau/passlib.v1"
)

func init() {
	// _ = os.Setenv("TZ", "America/Toronto")
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
	fmt.Println(dbdriver)
	fmt.Println(conn)

	_ = orm.RegisterDataBase("default", dbdriver, conn)
	_ = passlib.UseDefaults(passlib.DefaultsLatest)
}
func main() {
	orm.Debug = true
	app.InitBeego()
	beego.Run()
}
