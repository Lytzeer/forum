package forum

import (
	"fmt"
	fl "forum/Login"
	fr "forum/Register"
	"html/template"
	"io"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, err)
	}
	return
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/register.html"))
	namer := r.FormValue("name")
	mail := r.FormValue("mail-name")
	passwordr := r.FormValue("password-login")
	fr.Register(namer, mail, passwordr)
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, err)
	}
	return
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("mail-name")
	password := r.FormValue("password-login")
	fmt.Println(name, password)
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/login.html"))
	val := fl.Login(name, password)
	if val == 1 {
		fmt.Fprintf(io.Discard, "coucou")
		fmt.Println("coucou")
	}
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, err)
	}
	return
}
