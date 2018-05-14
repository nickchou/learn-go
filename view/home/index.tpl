<html>
<head>
	<title>index page</title>
</head>
<body>
	hi,{{.name}} <br/>
	email:{{.email}}<br/>
	hello world
	<hr/>
	学生列表:<br/>
	{{with .stus}}
	{{range .}}
		{{.Name}},{{.Sex}},{{.Age}} <br/>
	{{end}}
	{{end}}
</body>
</html>