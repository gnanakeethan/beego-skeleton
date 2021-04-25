package controllers

import (
	"backend/app/base/controllers"
)


type MainController struct {
	controllers.ActionController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
