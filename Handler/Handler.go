package forum

import (
	"fmt"
	fcr "forum/Create"
	fd "forum/Datas"
	fl "forum/Login"
	fr "forum/Register"
	ft "forum/Topic"
	"html/template"
	"net/http"
)

type displayerror struct {
	Leprobleme string
	Chargee    bool
}

var t displayerror
var cucu fd.User

var Topics []fd.Topic

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	Topics = ft.GetTopics()
	fmt.Println(Topics)
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/index.html"))
	err := tmpl.Execute(w, Topics)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, err)
	}
	return
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/chat.html"))
	create := r.FormValue("chat1")
	fcr.Create(create, cucu.User_name)
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
	confipass := r.FormValue("conf-password-login")
	fr.Register(namer, mail, passwordr, confipass)
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
	cucu.User_name = name
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/login.html"))
	val := fl.Login(name, password)
	if val == 1 {
		t.Leprobleme = "Invalid Username or Password"
		err := tmpl.Execute(w, t)
		if err != nil {
			fmt.Println(err)
			fmt.Fprintln(w, err)
		}
	} else {
		t.Leprobleme = "Login Successful"
		err := tmpl.Execute(w, t)
		if err != nil {
			fmt.Println(err)
			fmt.Fprintln(w, err)
		}
	}

	return
}
