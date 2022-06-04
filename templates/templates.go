package templates

import (
	"embed"
	"html/template"
)

//go:embed *.html
var templatesFS embed.FS

var Templates = template.Must(template.ParseFS(templatesFS, "*.html"))
