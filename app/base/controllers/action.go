package controllers

import (
	"html/template"

	beego "github.com/beego/beego/v2/server/web"
)

type ActionController struct {
	GenericController
}

func (actionController *ActionController) Prepare() {
	actionController.GenerateSession(nil)
	actionController.AdminLayout()
	actionController.Data["AppName"] = beego.AppConfig.DefaultString("appname", "Honey")
	actionController.Data["xsrfdata"] = template.HTML(actionController.XSRFFormHTML())
	actionController.Data["xsrftoken"] = actionController.XSRFToken()
}
