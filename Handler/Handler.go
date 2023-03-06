package forum

import (
	"fmt"
	fd "forum/Datas"
	ff "forum/Funcs"
	"html/template"
	"net/http"
	"strconv"
)

type displayerror struct {
	Leprobleme string
}

var t displayerror
var User fd.User
var Topics []fd.Topic
var Topic fd.Topic
var leboule bool

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if User.SignIn == true {
		token, err := r.Cookie("session")
		ff.CheckErr(err)
		fmt.Println(token.Value)
		User.Id, User.User_name, User.Email, User.Token = ff.CheckToken(token.Value)
		fmt.Println(User.Id, User.User_name, User.Email, User.Token, token.Value)
	}
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
		if User.Token == "" {
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
		} else {
			tmpl = template.Must(template.ParseFiles("./static/create.html"))
			title, message, categorie := r.FormValue("title"), r.FormValue("message"), r.FormValue("checkbox")
			fmt.Println(title, message, categorie)
			t.Leprobleme = ff.Create(message, User.User_name, title, "sport")
			if t.Leprobleme == "Topic created" {
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
		name, mail, password, confipass := r.FormValue("name"), r.FormValue("mail"), r.FormValue("password-register"), r.FormValue("conf-password-register")
		t.Leprobleme = ff.Register(name, mail, password, confipass)
		if t.Leprobleme == "Register successful" {
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
		}
		err := tmpl.Execute(w, t)
		ff.CheckErr(err)
		return
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		ff.Error404(w, r)
		return
	} else {
		name, password := r.FormValue("mail-name"), r.FormValue("password-login")
		User.User_name = name
		var tmpl *template.Template
		t.Leprobleme = ff.Login(name, password)
		if t.Leprobleme == "Login successful" {
			tmpl = template.Must(template.ParseFiles("./static/profile.html"))
			t.Leprobleme = ""
			User.Token = ff.SetCookie(w, r, User.User_name)
			err := tmpl.Execute(w, User)
			ff.CheckErr(err)
		} else {
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
			t.Leprobleme = ""
			err := tmpl.Execute(w, t)
			ff.CheckErr(err)
		}
		return
	}
}

func HandleProfile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/profile" {
		ff.Error404(w, r)
		return
	} else {
		if User.Token == "" {
			var tmpl *template.Template
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
			err := tmpl.Execute(w, nil)
			ff.CheckErr(err)
			return
		} else {
			var tmpl *template.Template
			tmpl = template.Must(template.ParseFiles("./static/profile.html"))
			err := tmpl.Execute(w, User)
			ff.CheckErr(err)
			return
		}
	}
}

func HandleAddComment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/addcomment" {
		ff.Error404(w, r)
		return
	} else {
		var tmpl *template.Template
		if User.Token == "" {
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
			err := tmpl.Execute(w, nil)
			ff.CheckErr(err)
			return
		} else {
			comment := r.FormValue("message")
			ff.AddComment(comment, User.User_name, Topic.TopicID, Topic.TopicAuthor)
			Topic = ff.GetOneTopics(Topic.TopicID)
			tmpl = template.Must(template.ParseFiles("./static/infos.html"))
			http.Redirect(w, r, "/infos?id="+strconv.Itoa(Topic.TopicID), 302)
			err := tmpl.Execute(w, Topic)
			ff.CheckErr(err)
			return
		}
	}
}

func HandleDeleteComment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/delcomment" {
		ff.Error404(w, r)
		return
	} else {
		var tmpl *template.Template
		if User.Token == "" {
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
			err := tmpl.Execute(w, nil)
			ff.CheckErr(err)
			return
		} else {
			id := r.FormValue("delete")
			idstr, _ := strconv.Atoi(id)
			ff.DeleteComment(User.User_name, idstr)
			Topic = ff.GetOneTopics(Topic.TopicID)
			tmpl = template.Must(template.ParseFiles("./static/infos.html"))
			http.Redirect(w, r, "/infos?id="+strconv.Itoa(Topic.TopicID), 302)
			err := tmpl.Execute(w, Topic)
			ff.CheckErr(err)
			return
		}
	}
}

func HandleModifyComment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/addcomment" {
		ff.Error404(w, r)
		return
	} else {
		var tmpl *template.Template
		if User.Token == "" {
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
			err := tmpl.Execute(w, nil)
			ff.CheckErr(err)
			return
		} else {
			comment := r.FormValue("message")
			ff.ModifyComment(comment, User.User_name, Topic.TopicID)
			tmpl = template.Must(template.ParseFiles("./static/infos.html"))
			err := tmpl.Execute(w, Topic)
			ff.CheckErr(err)
			return
		}
	}
}

func HandleDeleteTopic(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/delete" {
		ff.Error404(w, r)
		return
	} else {
		var tmpl *template.Template
		if User.Token == "" {
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
			err := tmpl.Execute(w, nil)
			ff.CheckErr(err)
			return
		} else {
			id := r.FormValue("delete")
			idstr, _ := strconv.Atoi(id)
			ff.DeleteTopic(User.User_name, idstr)
			Topics = ff.GetTopics()
			tmpl = template.Must(template.ParseFiles("./static/index.html"))
			err := tmpl.Execute(w, Topics)
			ff.CheckErr(err)
			return
		}
	}
}

func HandleModifyTopic(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/modifyt" {
		ff.Error404(w, r)
		return
	} else {
		var tmpl *template.Template
		if User.Token == "" {
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
			err := tmpl.Execute(w, nil)
			ff.CheckErr(err)
			return
		} else {
			comment := r.FormValue("message")
			ff.ModifyComment(comment, User.User_name, Topic.TopicID)
			tmpl = template.Must(template.ParseFiles("./static/index.html"))
			err := tmpl.Execute(w, Topic)
			ff.CheckErr(err)
			return
		}
	}
}

func HandleInfos(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/infos" {
		ff.Error404(w, r)
		return
	} else {
		id := r.FormValue("id")
		Topic.TopicID, _ = strconv.Atoi(id)
		Topic = ff.GetOneTopics(Topic.TopicID)
		Topic.Comments = ff.GetCommmentsOfTopic(Topic.TopicID)
		var tmpl *template.Template
		tmpl = template.Must(template.ParseFiles("./static/infos.html"))
		err := tmpl.Execute(w, Topic)
		ff.CheckErr(err)
		return
	}
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("e")
	if r.URL.Path != "/logout" {
		ff.Error404(w, r)
		return
	} else {
		User.User_name = ""
		User.Token = ""
		User.SignIn = true
		User.Id = 0
		ff.DeleteCookie(w, r)
		var tmpl *template.Template
		tmpl = template.Must(template.ParseFiles("./static/index.html"))
		err := tmpl.Execute(w, Topics)
		ff.CheckErr(err)
		http.Redirect(w, r, "/", 302)
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
		if User.Token == "" {
			var tmpl *template.Template
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
			err := tmpl.Execute(w, nil)
			ff.CheckErr(err)
			return
		} else {
			newemail, confemail, currentpassword, newpassword, confnewpassword, newusername := r.FormValue("newemail"), r.FormValue("confemail"), r.FormValue("currentpassword"), r.FormValue("newpassword"), r.FormValue("confnewpassword"), r.FormValue("newusername")
			fmt.Println(newemail, confemail, currentpassword, newpassword, confnewpassword, newusername)
			if newemail != "" && confemail != "" {
				ff.EditMail(User.Id, newemail, confemail)
			}
			if currentpassword != "" && newpassword != "" && confnewpassword != "" {
				ff.EditPassword(User.Id, currentpassword, newpassword, confnewpassword)
			}
			if newusername != "" {
				ff.EditUsername(User.Id, newusername)
			}
			var tmpl *template.Template
			tmpl = template.Must(template.ParseFiles("./static/editprofile.html"))
			err := tmpl.Execute(w, User)
			ff.CheckErr(err)
			return
		}
	}
}

func HandleNotif(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/notif" {
		ff.Error404(w, r)
		return
	} else {
		if User.Token == "" {
			var tmpl *template.Template
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
			err := tmpl.Execute(w, nil)
			ff.CheckErr(err)
			return
		} else {
			var Notifications []fd.Notif
			Notifications = ff.GetNotifs(User.User_name)
			var tmpl *template.Template
			tmpl = template.Must(template.ParseFiles("./static/notif.html"))
			err := tmpl.Execute(w, Notifications)
			ff.CheckErr(err)
			return
		}
	}
}

func HandleLike(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/like" {
		ff.Error404(w, r)
		return
	} else {
		if User.Token == "" {
			var tmpl *template.Template
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
			err := tmpl.Execute(w, nil)
			ff.CheckErr(err)
			return
		} else {
			id := r.FormValue("like")
			idint, _ := strconv.Atoi(id)
			fmt.Println(idint)
			if leboule == false {
				ff.Like(idint)
				leboule = true
			}
			var tmpl *template.Template
			Topic = ff.GetOneTopics(Topic.TopicID)
			tmpl = template.Must(template.ParseFiles("./static/infos.html"))
			http.Redirect(w, r, "/infos?id="+strconv.Itoa(Topic.TopicID), 302)
			err := tmpl.Execute(w, Topic)
			ff.CheckErr(err)
			return
		}
	}
}

func HandleDislike(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dislike" {
		ff.Error404(w, r)
		return
	} else {
		if User.Token == "" {
			var tmpl *template.Template
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
			err := tmpl.Execute(w, nil)
			ff.CheckErr(err)
			return
		} else {
			id := r.FormValue("dislike")
			idint, _ := strconv.Atoi(id)
			fmt.Println(idint)
			ff.Dislike(idint)
			var tmpl *template.Template
			Topic = ff.GetOneTopics(Topic.TopicID)
			tmpl = template.Must(template.ParseFiles("./static/infos.html"))
			http.Redirect(w, r, "/infos?id="+strconv.Itoa(Topic.TopicID), 302)
			err := tmpl.Execute(w, Topic)
			ff.CheckErr(err)
			return
		}
	}
}
