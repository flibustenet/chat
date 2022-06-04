package handlers

import (
	"net/http"

	"github.com/flibustenet/chat/templates"
)

func Index(w http.ResponseWriter, r *http.Request) {
	templates.Templates.ExecuteTemplate(w, "index.html", nil)
}
