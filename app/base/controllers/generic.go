package controllers

import (
	"html/template"

	"backend/models"

	beego "github.com/beego/beego/v2/server/web"
)

type GenericController struct {
	User *models.User
	beego.Controller
}

func (genericController *GenericController) Prepare() {
	genericController.GuestLayout()
	genericController.Data["AppName"] = beego.AppConfig.DefaultString("appname", "Honey")
	genericController.Data["xsrfdata"] = template.HTML(genericController.XSRFFormHTML())
	genericController.Data["xsrftoken"] = genericController.XSRFToken()
}

func (genericController *GenericController) GuestLayout() {
	genericController.Layout = "layout/guest.tpl"
	genericController.LayoutSections = make(map[string]string)
	genericController.LayoutSections["HtmlHead"] = "layout/guest/html_head.tpl"
	genericController.LayoutSections["Topbar"] = "layout/guest/topbar.tpl"
	genericController.LayoutSections["Scripts"] = "layout/guest/scripts.tpl"
	genericController.LayoutSections["Styles"] = "layout/guest/styles.tpl"
	genericController.LayoutSections["Sidebar"] = "layout/guest/sidebar.tpl"
}

func (genericController *GenericController) AdminLayout() {
	genericController.Layout = "layout/admin.tpl"
	genericController.LayoutSections = make(map[string]string)
	genericController.LayoutSections["HtmlHead"] = "layout/admin/html_head.tpl"
	genericController.LayoutSections["Topbar"] = "layout/admin/topbar.tpl"
	genericController.LayoutSections["Scripts"] = "layout/admin/scripts.tpl"
	genericController.LayoutSections["Styles"] = "layout/admin/styles.tpl"
	genericController.LayoutSections["Sidebar"] = "layout/admin/sidebar.tpl"
}

func (genericController *GenericController) GenerateSession(u *models.User) {
	if u != nil {
		genericController.SetSession("user_id", u.Id)
		genericController.User = u
		genericController.Data["User"] = u
		return
	} else {
		genericController.StartSession()
		userInterface := genericController.GetSession("user_id")
		if uObj, ok := userInterface.(string); ok {
			genericController.User, _ = models.GetUserById(uObj)
			genericController.Data["User"] = genericController.User
			return
		} else if genericController.Ctx.Request.RequestURI != "/auth/login" {
			genericController.Redirect("/auth/login", 301)
			return
		}
	}
}
