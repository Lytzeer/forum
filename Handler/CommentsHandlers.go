package forum

import (
	ff "forum/Funcs"
	"html/template"
	"net/http"
	"strconv"
)

func HandleAddComment(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
	if r.URL.Path != "/addcomment" {
		ff.Error404(w, r)
		return
	} else {
		if User.Token == "" {
			http.Redirect(w, r, "/login", 302)
		} else {
			comment := r.FormValue("message")
			ff.AddComment(comment, User.User_name, Topic.TopicID, Topic.TopicAuthor)
			http.Redirect(w, r, "/infos?id="+strconv.Itoa(Topic.TopicID), 302)
			return
		}
	}
}

func HandleDeleteComment(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
	if r.URL.Path != "/delcomment" {
		ff.Error404(w, r)
		return
	} else {
		if User.Token == "" {
			http.Redirect(w, r, "/login", 302)
		} else {
			id := r.FormValue("delete")
			idstr, _ := strconv.Atoi(id)
			ff.DeleteComment(User.User_name, idstr)
			http.Redirect(w, r, "/infos?id="+strconv.Itoa(Topic.TopicID), 302)
			return
		}
	}
}

func HandleModifyComment(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
	if r.URL.Path != "/modifycomment" {
		ff.Error404(w, r)
		return
	} else {
		var tmpl *template.Template
		if User.Token == "" {
			http.Redirect(w, r, "/login", 302)
		} else {
			commentId, message := r.FormValue("modify"), r.FormValue("message")
			if modifycommentid == 0 {
				modifycommentid, _ = strconv.Atoi(commentId)
			}
			commentIdstr, _ := strconv.Atoi(commentId)
			if commentIdstr == 0 {
				ff.ModifyComment(message, User.User_name, modifycommentid)
				http.Redirect(w, r, "/infos?id="+strconv.Itoa(Topic.TopicID), 302)
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

func HandleLike(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
	if r.URL.Path != "/like" {
		ff.Error404(w, r)
		return
	} else {
		if User.Token == "" {
			http.Redirect(w, r, "/login", 302)
		} else {
			id := r.FormValue("like")
			idint, _ := strconv.Atoi(id)
			ff.Like(idint, User.User_name, Comment.CreatorName)
			http.Redirect(w, r, "/infos?id="+strconv.Itoa(Topic.TopicID), 302)
			return
		}
	}
}

func HandleDislike(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
	if r.URL.Path != "/dislike" {
		ff.Error404(w, r)
		return
	} else {
		if User.Token == "" {
			http.Redirect(w, r, "/login", 302)
		} else {
			id := r.FormValue("dislike")
			idint, _ := strconv.Atoi(id)
			ff.Dislike(idint, User.User_name, Comment.CreatorName)
			http.Redirect(w, r, "/infos?id="+strconv.Itoa(Topic.TopicID), 302)
			return
		}
	}
}
