package main

import (
	"backend/app"
	_ "backend/routers"

	_ "github.com/lib/pq"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	orm.Debug = true
	app.InitBeego()
	beego.Run()
}
