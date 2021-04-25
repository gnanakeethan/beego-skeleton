<!DOCTYPE html>
<html>
<head>
<title>{{.Title}}</title>
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta name="_xsrf" content="{{.xsrftoken}}" />
{{.HtmlHead}}
{{.Styles}}
</head>
<body>
{{.Topbar}}
{{.SideBar}}
{{.LayoutContent}}
{{.Scripts}}
</body>
</html>