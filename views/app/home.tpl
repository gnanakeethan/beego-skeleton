{{.User.VisibleName}}
{{ if hasRole .User "'superadmin'" "'customer'"}}
	Super
{{end}}
{{ if hasPermission .User "soemthing" "of"}}
	Is Super Admin
{{ end}}