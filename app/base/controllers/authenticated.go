package controllers

import (
	"html/template"

	beego "github.com/beego/beego/v2/server/web"
)

type AuthenticatedController struct {
	GenericController
}

func (authenticatedController *AuthenticatedController) Prepare() {
	authenticatedController.AdminLayout()
	authenticatedController.GenerateSession(nil)
	authenticatedController.Data["AppName"] = beego.AppConfig.DefaultString("appname", "Honey")
	authenticatedController.Data["xsrfdata"] = template.HTML(authenticatedController.XSRFFormHTML())
	authenticatedController.Data["xsrftoken"] = authenticatedController.XSRFToken()
}
