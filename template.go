package webi

var template_string = `
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" type="text/css" href="/style.css" media="screen" />
	</head>
	<body>
		<div id="holder">
			<div id="header">
				<h1>{{.Header}}</h1>
			</div>
			<div id="content">
				{{range .Content}}
				<p>{{.}}</p>
				{{end}}
			</div>
			<div id="footer">
			</div>
		</div>
	</body>
</html>`