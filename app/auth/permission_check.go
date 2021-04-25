package auth

import (
	"backend/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
)

func HasRole(user *models.User, role ...string) bool {
	return user.HasRole(role...)
}

func HasPermission(user *models.User, permissionString, method string) bool {
	o := orm.NewOrm()
	superadmin := user.HasRole("'superadmin'")
	if superadmin {
		return true
	} else if user != nil && !superadmin {
		var permissions []*models.Permission
		customquery := "select p.* from users u inner join role_user ru on ru.user_id=u.id " +
			"inner join roles r on ru.role_id=r.id " +
			"inner join permission_role rp on rp.role_id=r.id " +
			"inner join permissions p on p.id=rp.permission_id " +
			"where UPPER(p.route)= UPPER(?) and UPPER(p.method)= UPPER(?) and u.id =?"
		logs.Debug("RAN IN 2")
		if count, err := o.Raw(customquery, permissionString, method, user.Id).QueryRows(&permissions); err == nil && count > 0 {
			logs.Debug(count)
			return true
		}
	}
	return false
}

func IsLoggedIn(this *context.Context) {
	userInterface := this.Input.Session("user_id")
	if userId, ok := userInterface.(string); ok {
		if user, err := models.GetUserById(userId); err == nil && user != nil {
			logs.Info(user)
			return
		}
	}
	this.Redirect(301, "/auth/login")
}

func IsAuthorized(this *context.Context) {
	userInterface := this.Input.Session("user_id")
	if userId, ok := userInterface.(string); ok {
		logs.Info(userId)
		if user, err := models.GetUserById(userId); err == nil && user != nil && HasPermission(user, this.Request.RequestURI, this.Request.Method) {
			return
		}
	}
	this.Redirect(301, "/auth/login")
}
