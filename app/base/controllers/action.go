package controllers

import (
	"html/template"

	beego "github.com/beego/beego/v2/server/web"
)

type ActionController struct {
	GenericController
}

func (actionController *ActionController) Prepare() {
	actionController.Layout = "layout/admin.tpl"
	actionController.LayoutSections = make(map[string]string)
	actionController.LayoutSections["HtmlHead"] = "layout/html_head.tpl"
	actionController.LayoutSections["Topbar"] = "layout/topbar.tpl"
	actionController.LayoutSections["Scripts"] = "layout/scripts.tpl"
	actionController.LayoutSections["Styles"] = "layout/styles.tpl"
	actionController.LayoutSections["Sidebar"] = "layout/sidebar.tpl"
	actionController.Data["AppName"] = beego.AppConfig.DefaultString("appname","Honey")
	actionController.Data["xsrfdata"] = template.HTML(actionController.XSRFFormHTML())
	actionController.Data["xsrftoken"] = actionController.XSRFToken()
}
