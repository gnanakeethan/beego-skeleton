package controllers

import (
	"html/template"
	"strings"

	"backend/app/lang"
	"backend/models"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

type GenericController struct {
	User *models.User
	beego.Controller
	i18n.Locale
}

func (genericController *GenericController) Prepare() {
	genericController.setLangVer()
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
		} else if !strings.Contains(genericController.Ctx.Request.RequestURI, "/auth/login") {
			genericController.Redirect("/auth/login?redirect_back="+genericController.Ctx.Request.RequestURI, 302)
			return
		}
	}
}

// setLangVer sets site language version.
func (this *GenericController) setLangVer() bool {
	isNeedRedir := false
	hasCookie := false


	// 1. Check URL arguments.
	preferredLang := this.GetString("preferredLang", "")

	// 2. Get language information from cookies.
	if len(preferredLang) == 0 {
		preferredLang = this.Ctx.GetCookie("preferredLang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	// Check again in case someone modify on purpose.
	if !i18n.IsExist(preferredLang) {
		preferredLang = ""
		isNeedRedir = false
		hasCookie = false
	}

	// 3. Get language information from 'Accept-Language'.
	if len(preferredLang) == 0 {
		al := this.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				preferredLang = al
			}
		}
	}

	// 4. Default language is English.
	if len(preferredLang) == 0 {
		preferredLang = "ta-LK"
		isNeedRedir = false
	}

	curLang := lang.LangType{
		Lang: preferredLang,
	}

	// Save language information in cookies.
	if !hasCookie {
		this.Ctx.SetCookie("preferredLang", curLang.Lang, 1<<31-1, "/")
	}

	restLangs := make([]*lang.LangType, 0, len(lang.LangTypes)-1)
	for _, v := range lang.LangTypes {
		if preferredLang != v.Lang {
			restLangs = append(restLangs, v)
		} else {
			curLang.Name = v.Name
		}
	}
	// Set language properties.
	this.Lang = preferredLang
	this.Data["Lang"] = curLang.Lang
	this.Data["CurLang"] = curLang.Name
	this.Data["RestLangs"] = restLangs
	return isNeedRedir
}
