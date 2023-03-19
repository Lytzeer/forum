package forum

import (
	ff "forum/Funcs"
	"html/template"
	"net/http"
	"strconv"
)

func HandleDeleteTopic(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
	if r.URL.Path != "/delete" {
		ff.Error404(w, r)
		return
	} else {
		var tmpl *template.Template
		if User.Token == "" {
			http.Redirect(w, r, "/login", 302)
		} else {
			topicid := r.FormValue("delete")
			if modifytopicid == 0 {
				modifytopicid, _ = strconv.Atoi(topicid)
			}
			idstr, _ := strconv.Atoi(topicid)
			ff.DeleteTopic(User.User_name, idstr)
			T.Topics = ff.GetTopics()
			tmpl = template.Must(template.ParseFiles("./static/index.html"))
			err := tmpl.Execute(w, T)
			ff.CheckErr(err)
			return
		}
	}
}

func HandleModifyTopic(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
	if r.URL.Path != "/modifyt" {
		ff.Error404(w, r)
		return
	} else {
		var tmpl *template.Template
		if User.Token == "" {
			http.Redirect(w, r, "/login", 302)
		} else {
			comment, title, topicid := r.FormValue("message"), r.FormValue("title"), r.FormValue("modify")
			if modifytopicid == 0 {
				modifytopicid, _ = strconv.Atoi(topicid)
			}
			topicidstr, _ := strconv.Atoi(topicid)
			if topicidstr == 0 {
				ff.ModifyTopic(title, comment, User.User_name, modifytopicid)
				http.Redirect(w, r, "/", 302)
				return
			} else {
				T2.Topic = ff.GetOneTopics(modifytopicid)
				tmpl = template.Must(template.ParseFiles("./static/edittopic.html"))
				err := tmpl.Execute(w, T2)
				ff.CheckErr(err)
				return
			}

		}
	}
}

func HandleInfos(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
	if r.URL.Path != "/infos" {
		ff.Error404(w, r)
		return
	} else {
		id := r.FormValue("id")
		Topic.TopicID, _ = strconv.Atoi(id)
		Topic, Topic.Comments = ff.GetOneTopics(Topic.TopicID), ff.GetCommmentsOfTopic(Topic.TopicID)
		T2.Topic, T2.User = Topic, User
		var tmpl *template.Template
		tmpl = template.Must(template.ParseFiles("./static/infos.html"))
		err := tmpl.Execute(w, T2)
		ff.CheckErr(err)
		return
	}
}

func HandleLikeTopic(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
	if r.URL.Path != "/liketopic" {
		ff.Error404(w, r)
		return
	} else {
		if User.Token == "" {
			http.Redirect(w, r, "/login", 302)
		} else {
			ff.LikeTopic(T2.Topic.TopicID, User.User_name, Topic.TopicAuthor)
			http.Redirect(w, r, "/infos?id="+strconv.Itoa(T2.Topic.TopicID), 302)
			return
		}
	}
}

func HandleDislikeTopic(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
	if r.URL.Path != "/disliketopic" {
		ff.Error404(w, r)
		return
	} else {
		if User.Token == "" {
			http.Redirect(w, r, "/login", 302)
		} else {
			ff.DislikeTopic(T2.Topic.TopicID, User.User_name, Topic.TopicAuthor)
			http.Redirect(w, r, "/infos?id="+strconv.Itoa(T2.Topic.TopicID), 302)
			return
		}
	}
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
	if r.URL.Path != "/create" {
		ff.Error404(w, r)
		return
	} else {
		var tmpl *template.Template
		if User.Token == "" {
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
		} else {
			tmpl = template.Must(template.ParseFiles("./static/create.html"))
			title, message, categorie := r.FormValue("title"), r.FormValue("message"), r.FormValue("categories")
			Error.Error = ff.Create(message, User.User_name, title, categorie)
			if Error.Error == "Topic created" {
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
