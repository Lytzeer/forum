package forum

import (
	"html/template"
	"net/http"
)

func Error404(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/error404.html"))
	tmpl.Execute(w, nil)
	return
}
