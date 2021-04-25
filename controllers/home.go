package controllers

import (
	"backend/app/base/controllers"
)

type HomeController struct {
	controllers.GenericController
}

func (c *HomeController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.GenerateSession(nil)
	c.TplName = "index.tpl"
}
