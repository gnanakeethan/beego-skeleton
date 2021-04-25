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

func (actionController *GenericController) Prepare() {
	actionController.Layout = "layout/guest.tpl"
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


func (this *GenericController) GenerateSession(u *models.User) {
	if u != nil {
		this.SetSession("user_id", u.Id)
		this.User = u
		this.Data["User"] = u
		return
	} else {
		this.StartSession()
		userInterface := this.GetSession("user_id")
		if uObj, ok := userInterface.(string); ok {
			this.User, _ = models.GetUserById(uObj)
			this.Data["User"] = this.User
			return
		} else if this.Ctx.Request.RequestURI != "/auth/login" {
			this.Redirect("/auth/login", 301)
			return
		}
	}
}
