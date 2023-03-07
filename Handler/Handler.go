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

var modifycommentid int
var modifytopicid int

var T fd.Topics
var T2 fd.TopicInfos

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
		T.Topics = Topics
		T.User = User
		var tmpl *template.Template
		tmpl = template.Must(template.ParseFiles("./static/index.html"))
		err := tmpl.Execute(w, T)
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
				T.Topics = ff.GetTopics()
				T.User = User
				err := tmpl.Execute(w, T)
				ff.CheckErr(err)
				return
			} else {
				tmpl = template.Must(template.ParseFiles("./static/create.html"))
				err := tmpl.Execute(w, nil)
				ff.CheckErr(err)
				return
			}
		}
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
			T2.Topic = Topic
			T2.User = User
			tmpl = template.Must(template.ParseFiles("./static/infos.html"))
			http.Redirect(w, r, "/infos?id="+strconv.Itoa(Topic.TopicID), 302)
			err := tmpl.Execute(w, T2)
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
			T2.Topic = Topic
			T2.User = User
			tmpl = template.Must(template.ParseFiles("./static/infos.html"))
			http.Redirect(w, r, "/infos?id="+strconv.Itoa(Topic.TopicID), 302)
			err := tmpl.Execute(w, T2)
			ff.CheckErr(err)
			return
		}
	}
}

func HandleModifyComment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/modifycomment" {
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
			commentId := r.FormValue("modify")
			if modifycommentid == 0 {
				modifycommentid, _ = strconv.Atoi(commentId)
			}
			commentIdstr, _ := strconv.Atoi(commentId)
			message := r.FormValue("message")
			if commentIdstr == 0 {
				ff.ModifyComment(message, User.User_name, modifycommentid)
				T2.Topic = ff.GetOneTopics(modifycommentid)
				tmpl = template.Must(template.ParseFiles("./static/infos.html"))
				http.Redirect(w, r, "/infos?id="+strconv.Itoa(Topic.TopicID), 302)
				err := tmpl.Execute(w, T2)
				ff.CheckErr(err)
				return
			} else {
				Comment := ff.GetOneComment(modifycommentid)
				tmpl = template.Must(template.ParseFiles("./static/editcomment.html"))
				err := tmpl.Execute(w, Comment)
				ff.CheckErr(err)
				return
			}
		}
	}
}

// pute
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
			topicid := r.FormValue("delete")
			if modifytopicid == 0 {
				modifytopicid, _ = strconv.Atoi(topicid)
			}
			idstr, _ := strconv.Atoi(topicid)
			ff.DeleteTopic(User.User_name, idstr)
			Topics = ff.GetTopics()
			T.Topics = Topics
			tmpl = template.Must(template.ParseFiles("./static/index.html"))
			err := tmpl.Execute(w, T)
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
			title := r.FormValue("title")
			fmt.Println("comment", comment, "title", title)
			topicid := r.FormValue("modify")
			if modifytopicid == 0 {
				modifytopicid, _ = strconv.Atoi(topicid)
			}
			topicidstr, _ := strconv.Atoi(topicid)
			if topicidstr == 0 {
				ff.ModifyTopic(title, comment, User.User_name, modifytopicid)
				T.Topics = ff.GetTopics()
				tmpl = template.Must(template.ParseFiles("./static/index.html"))
				http.Redirect(w, r, "/", 302)
				err := tmpl.Execute(w, T)
				ff.CheckErr(err)
				return
			} else {
				Topic = ff.GetOneTopics(modifytopicid)
				T2.Topic = Topic
				tmpl = template.Must(template.ParseFiles("./static/edittopic.html"))
				err := tmpl.Execute(w, T2)
				ff.CheckErr(err)
				return
			}

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
		T2.Topic = Topic
		T2.User = User
		var tmpl *template.Template
		tmpl = template.Must(template.ParseFiles("./static/infos.html"))
		err := tmpl.Execute(w, T2)
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
		T.User = User
		ff.DeleteCookie(w, r)
		var tmpl *template.Template
		tmpl = template.Must(template.ParseFiles("./static/index.html"))
		err := tmpl.Execute(w, T)
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
			Topic.Comments = ff.GetCommmentsOfTopic(Topic.TopicID)
			T2.Topic = Topic
			T2.User = User
			tmpl = template.Must(template.ParseFiles("./static/infos.html"))
			http.Redirect(w, r, "/infos?id="+strconv.Itoa(Topic.TopicID), 302)
			err := tmpl.Execute(w, T2)
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
			Topic.Comments = ff.GetCommmentsOfTopic(Topic.TopicID)
			T2.Topic = Topic
			T2.User = User
			tmpl = template.Must(template.ParseFiles("./static/infos.html"))
			http.Redirect(w, r, "/infos?id="+strconv.Itoa(Topic.TopicID), 302)
			err := tmpl.Execute(w, T2)
			ff.CheckErr(err)
			return
		}
	}
}

func HandleLikeTopic(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/liketopic" {
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
			ff.LikeTopic(T2.Topic.TopicID)
			T2.Topic = ff.GetOneTopics(T2.Topic.TopicID)
			T2.Topic.Comments = ff.GetCommmentsOfTopic(T2.Topic.TopicID)
			T2.User = User
			var tmpl *template.Template
			tmpl = template.Must(template.ParseFiles("./static/infos.html"))
			http.Redirect(w, r, "/infos?id="+strconv.Itoa(T2.Topic.TopicID), 302)
			err := tmpl.Execute(w, T2)
			ff.CheckErr(err)
			return
		}
	}
}

func HandleDislikeTopic(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/disliketopic" {
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
			ff.DislikeTopic(T2.Topic.TopicID)
			T2.Topic = ff.GetOneTopics(T2.Topic.TopicID)
			T2.Topic.Comments = ff.GetCommmentsOfTopic(T2.Topic.TopicID)
			T2.User = User
			var tmpl *template.Template
			tmpl = template.Must(template.ParseFiles("./static/infos.html"))
			http.Redirect(w, r, "/infos?id="+strconv.Itoa(T2.Topic.TopicID), 302)
			err := tmpl.Execute(w, T2)
			ff.CheckErr(err)
			return
		}
	}
}

func HandleFilters(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/filters" {
		ff.Error404(w, r)
		return
	} else {
		filter := r.FormValue("filter")
		T.Topics = ff.GetFiltred(filter)
		var tmpl *template.Template
		tmpl = template.Must(template.ParseFiles("./static/index.html"))
		err := tmpl.Execute(w, T)
		ff.CheckErr(err)
		return
	}
}
