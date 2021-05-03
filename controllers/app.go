package controllers

import (
	"backend/app/base/controllers"
)

type AppController struct {
	controllers.AuthenticatedController
}
func (c *AppController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// Get
// @Title Get home page
// @Success 201 {int} html
// @router / [get]
func (c *AppController) Get() {
	c.Data["Title"] = "App"
	c.TplName = "app/home.tpl"
}
