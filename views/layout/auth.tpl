<!DOCTYPE html>
<html>
<head>
	<title>{{.AppName}} | {{.Title}}</title>
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<link rel="stylesheet" href="/static/css/app.css">

	<meta name="_xsrf" content="{{.xsrftoken}}" />
    {{.HtmlHead}}
    {{.Styles}}
</head>
<body>
{{.LayoutContent}}
{{.Scripts}}
</body>
</html>