// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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
<style>
	body {
		background-image: linear-gradient(-45deg, white, yellow);
	}
	h1 {
		color: #2f4f4f;
		font-size: 3em;
		padding-top: 1.1em;
	}
</style>
</head>
<body>
	<center>
		<h1>
			{{.UUID}}
		</h1>
	</center>
</body>
</html>
`)
	if err != nil {
		log.Fatal("failed to render template:", err)
	}
}
