package main

import (
	"backend/app"
	_ "backend/routers"

	_ "github.com/lib/pq"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	app.InitBeego()
	beego.Run()
}
