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
	if r.URL.Path != "/" {
		ff.Error404(w, r)
		return
	} else {
		Topics = ff.GetTopics()
		var tmpl *template.Template
		tmpl = template.Must(template.ParseFiles("./static/index.html"))
		err := tmpl.Execute(w, Topics)
		ff.CheckErr(err)
		return
	}
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/create" {
		ff.Error404(w, r)
		return
	} else {
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
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		ff.Error404(w, r)
		return
	} else {
		var tmpl *template.Template
		tmpl = template.Must(template.ParseFiles("./static/register.html"))
		namer := r.FormValue("name")
		mail := r.FormValue("mail")
		passwordr := r.FormValue("password-register")
		confipass := r.FormValue("conf-password-register")
		t.Leprobleme = ff.Register(namer, mail, passwordr, confipass)
		if t.Leprobleme == "Register successful" {
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
		}
		err := tmpl.Execute(w, t)
		ff.CheckErr(err)
		return
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// a, err := r.Cookie("session")
	// ff.CheckErr(err)
	// fmt.Println(a)
	if r.URL.Path != "/login" {
		ff.Error404(w, r)
		return
	} else {
		name := r.FormValue("mail-name")
		password := r.FormValue("password-login")
		cucu.User_name = name
		var tmpl *template.Template
		t.Leprobleme = ff.Login(name, password)
		if t.Leprobleme == "Login successful" {
			tmpl = template.Must(template.ParseFiles("./static/profile.html"))
			t.Leprobleme = ""
			cucu.Token = ff.SetCookie(w, r, cucu.User_name)
		} else {
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
			t.Leprobleme = ""
		}
		err := tmpl.Execute(w, t)
		ff.CheckErr(err)
		return
	}
}

func HandleProfile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/profile" {
		ff.Error404(w, r)
		return
	} else {
		var tmpl *template.Template
		tmpl = template.Must(template.ParseFiles("./static/profile.html"))
		err := tmpl.Execute(w, nil)
		ff.CheckErr(err)
		return
	}
	// if cucu.User_name == "" {
	// 	tmpl = template.Must(template.ParseFiles("./static/login.html"))
	// } else {

	// }
}

func HandleComment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/comment" {
		ff.Error404(w, r)
		return
	} else {
		var tmpl *template.Template
		if cucu.User_name == "" {
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
		} else {
			comment := r.FormValue("comment")
			ff.AddComment(comment, cucu.User_name, 0)
			tmpl = template.Must(template.ParseFiles("./static/infos.html"))
			err := tmpl.Execute(w, nil)
			ff.CheckErr(err)
			return
		}
	}
}

func HandleInfos(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/info" {
		ff.Error404(w, r)
		return
	} else {
		var tmpl *template.Template
		tmpl = template.Must(template.ParseFiles("./static/infos.html"))
		err := tmpl.Execute(w, nil)
		ff.CheckErr(err)
		return
	}
	// TopComment = ff.GetOneTopics(1)
	// get id of topic
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" {
		ff.Error404(w, r)
		return
	} else {
		var tmpl *template.Template
		tmpl = template.Must(template.ParseFiles("./static/index.html"))
		cucu.User_name = ""
		cucu.Token = ""
		ff.DeleteCookie(w, r)
		err := tmpl.Execute(w, nil)
		ff.CheckErr(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func HandleTest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/test" {
		ff.Error404(w, r)
		return
	} else {
		var tmpl *template.Template
		tmpl = template.Must(template.ParseFiles("./static/test.html"))
		err := tmpl.Execute(w, nil)
		ff.CheckErr(err)
		return
	}
}

func HandleEditProfile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/editprofile" {
		ff.Error404(w, r)
		return
	} else {
		var tmpl *template.Template
		tmpl = template.Must(template.ParseFiles("./static/editprofile.html"))
		err := tmpl.Execute(w, nil)
		ff.CheckErr(err)
		return
	}
}
