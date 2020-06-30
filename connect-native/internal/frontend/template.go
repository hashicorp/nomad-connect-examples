package frontend

import (
	"html/template"
	"log"
)

var indexHTML *template.Template

func init() {
	var err error
	indexHTML, err = template.New("index").Parse(`
<html>
<head>
</head>
<body>
  <center>uuid: {{.UUID}}</center>
</body>
</html>
`)
	if err != nil {
		log.Fatal("failed to render template:", err)
	}
}
