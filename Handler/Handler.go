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
	fmt.Println("cucu", cucu, cucu.User_name)
	// cucu.User_name = ""
	// cucu.Email = ""
	fmt.Println("cucu", cucu, cucu.User_name)
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
	var a string
	if cucu.User_name == "" {
		tmpl = template.Must(template.ParseFiles("./static/login.html"))
	} else {
		tmpl = template.Must(template.ParseFiles("./static/create.html"))
		title := r.FormValue("title")
		message := r.FormValue("content")
		categorie := r.FormValue("checkbox")
		fmt.Println("title", title, "message", message, "categorie", categorie, "user", cucu.User_name)

		a := fcr.Create(message, cucu.User_name, title, categorie)
		fmt.Println(a)
	}
	if a == "Topic created" {
		tmpl = template.Must(template.ParseFiles("./static/index.html"))
	}

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
	t.Leprobleme = fr.Register(namer, mail, passwordr, confipass)
	if t.Leprobleme == "Register successful" {
		tmpl = template.Must(template.ParseFiles("./static/login.html"))
	}
	err := tmpl.Execute(w, t)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, err)
	}
	return
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("mail-name")
	password := r.FormValue("password-login")
	fmt.Println("name", name, "password", password)
	cucu.User_name = name
	fmt.Println("name", cucu.User_name, "password", password)
	var tmpl *template.Template
	t.Leprobleme = fl.Login(name, password)
	if t.Leprobleme == "Login successful" {
		tmpl = template.Must(template.ParseFiles("./static/profile.html"))
		t.Leprobleme = ""
	} else {
		tmpl = template.Must(template.ParseFiles("./static/login.html"))
		t.Leprobleme = ""
	}
	err := tmpl.Execute(w, t)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, err)
	}
	return
}

func HandleProfile(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	// if cucu.User_name == "" {
	// 	tmpl = template.Must(template.ParseFiles("./static/login.html"))
	// } else {
	tmpl = template.Must(template.ParseFiles("./static/profile.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, err)
	}
	return
	// }
}

func HandleInfos(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	// if cucu.User_name == "" {
	// 	tmpl = template.Must(template.ParseFiles("./static/login.html"))
	// } else {
	tmpl = template.Must(template.ParseFiles("./static/infos.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, err)
	}
	return
	// }
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/index.html"))
	cucu.User_name = ""
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
