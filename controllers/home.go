package controllers

import (
	"backend/app/base/controllers"
)

type HomeController struct {
	controllers.GenericController
}

func (c *HomeController) URLMapping() {
	c.Mapping("Get", c.Get)
}


// Get
// @Title Get home page
// @Success 201 {int} html
// @router / [get]
func (c *HomeController) Get() {
	c.Data["Title"] = "App"
	c.TplName = "index.tpl"
}
