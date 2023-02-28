package forum

import (
	"fmt"
	fd "forum/Datas"
	ff "forum/Funcs"
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
var TopComment []fd.Topic

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	Topics = ff.GetTopics()
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/index.html"))
	err := tmpl.Execute(w, Topics)
	ff.CheckErr(err)
	return
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	if cucu.User_name == "" {
		tmpl = template.Must(template.ParseFiles("./static/login.html"))
	} else {
		var a string
		tmpl = template.Must(template.ParseFiles("./static/create.html"))
		title := r.FormValue("title")
		message := r.FormValue("content")
		categorie := r.FormValue("checkbox")
		a = ff.Create(message, cucu.User_name, title, categorie)
		fmt.Println(a)
		if a == "Topic created" {
			tmpl = template.Must(template.ParseFiles("./static/index.html"))
		} else {
			tmpl = template.Must(template.ParseFiles("./static/create.html"))
		}
	}
	err := tmpl.Execute(w, nil)
	ff.CheckErr(err)
	return
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/register.html"))
	namer := r.FormValue("name")
	mail := r.FormValue("mail-name")
	passwordr := r.FormValue("password-login")
	confipass := r.FormValue("conf-password-login")
	t.Leprobleme = ff.Register(namer, mail, passwordr, confipass)
	if t.Leprobleme == "Register successful" {
		tmpl = template.Must(template.ParseFiles("./static/login.html"))
	}
	err := tmpl.Execute(w, t)
	ff.CheckErr(err)
	return
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("mail-name")
	password := r.FormValue("password-login")
	fmt.Println("name", name, "password", password)
	cucu.User_name = name
	fmt.Println("name", cucu.User_name, "password", password)
	var tmpl *template.Template
	t.Leprobleme = ff.Login(name, password)
	if t.Leprobleme == "Login successful" {
		tmpl = template.Must(template.ParseFiles("./static/profile.html"))
		t.Leprobleme = ""
	} else {
		tmpl = template.Must(template.ParseFiles("./static/login.html"))
		t.Leprobleme = ""
	}
	err := tmpl.Execute(w, t)
	ff.CheckErr(err)
	return
}

func HandleProfile(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	// if cucu.User_name == "" {
	// 	tmpl = template.Must(template.ParseFiles("./static/login.html"))
	// } else {
	tmpl = template.Must(template.ParseFiles("./static/profile.html"))
	err := tmpl.Execute(w, nil)
	ff.CheckErr(err)
	return
	// }
}

func HandleComment(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	if cucu.User_name == "" {
		tmpl = template.Must(template.ParseFiles("./static/login.html"))
	} else {
		tmpl = template.Must(template.ParseFiles("./static/infos.html"))
		err := tmpl.Execute(w, nil)
		ff.CheckErr(err)
		return
	}
}

func HandleInfos(w http.ResponseWriter, r *http.Request) {
	// TopComment = ff.GetOneTopics(1)
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/infos.html"))
	err := tmpl.Execute(w, nil)
	ff.CheckErr(err)
	return
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/index.html"))
	cucu.User_name = ""
	err := tmpl.Execute(w, nil)
	ff.CheckErr(err)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
