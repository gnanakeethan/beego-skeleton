package controllers

import (
	"html/template"

	"backend/app/base/controllers"
	"backend/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"gopkg.in/hlandau/passlib.v1"
)

type AuthController struct {
	controllers.GenericController
}

// Prepares the login interfaces
func (authController *AuthController) Prepare() {
	authController.Layout = "layout/guest.tpl"
	authController.LayoutSections = make(map[string]string)
	authController.LayoutSections["HtmlHead"] = "layout/guest/html_head.tpl"
	authController.LayoutSections["Scripts"] = "layout/guest/scripts.tpl"
	authController.LayoutSections["Styles"] = "layout/guest/styles.tpl"
	authController.Data["AppName"] = beego.AppConfig.DefaultString("appname", "Honey")
	authController.Data["xsrfdata"] = template.HTML(authController.XSRFFormHTML())
	authController.Data["xsrftoken"] = authController.XSRFToken()
}

func (authController *AuthController) Logout() {
	authController.DestroySession()
	authController.Redirect("/", 302)
}

func (authController *AuthController) Login() {
	authController.GenerateSession(nil)
	if authController.User != nil {
		authController.Redirect("/", 302)
	}
	authController.Data["Title"] = "Login"
	authController.Data["RedirectBack"] = authController.GetString("redirect_back", beego.AppConfig.DefaultString("login_redirect", ""))
	authController.TplName = "honeypot/auth/login.tpl"
}

func (authController *AuthController) LoginPost() {
	redirect := authController.GetString("redirect_back", beego.AppConfig.DefaultString("login_redirect", ""))
	authController.Data["Title"] = "Login"
	authController.TplName = "honeypot/auth/login.tpl"
	username := authController.GetString("username")
	password := authController.GetString("password")
	o := orm.NewOrm()
	user := models.User{Username: username}
	err := o.Read(&user, "username")
	pass := false
	if err == nil {
		logs.Info(user.Password)
		if _, err := passlib.Verify(password, user.Password); err == nil {
			pass = true
			goto LOGIN_SUCCESS
		}
		authController.Data["Flash"] = map[string]interface{}{
			"errored":    true,
			"error_type": "Validation Error",
			"fields":     map[string]string{"username": "invalid username / password combination"}}
		return
	}

LOGIN_SUCCESS:
	if pass {
		authController.GenerateSession(&user)
		if authController.User != nil {
			authController.Redirect(redirect, 302)
			return
		}
	}
	return
}
