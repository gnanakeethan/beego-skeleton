package main

import (
	"backend/app"
	_ "backend/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	app.InitBeego()
	beego.Run()
}
