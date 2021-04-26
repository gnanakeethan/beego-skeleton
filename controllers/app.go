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

// Post ...
// @Title Post
// @Success 201 {int} string
// @router / [get]
func (c *AppController) Get() {
	c.Data["Title"] = "App"
	c.TplName = "app/home.tpl"
}
