package forum

import (
	fd "forum/Datas"
	ff "forum/Funcs"
	"html/template"
	"net/http"
)

var Error fd.DisplayError
var User fd.User
var Topics []fd.Topic
var Topic fd.Topic
var Comment fd.Comment

var modifycommentid int
var modifytopicid int

var T fd.Topics
var T2 fd.TopicInfos

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
	// if User.SignIn == false {
	// 	token, _ := r.Cookie("session")
	// 	//User.Id, User.User_name, User.Email, User.Token = ff.CheckToken(token.Value)
	// 	fmt.Println(token.Value == User.Token)
	// 	fmt.Println(User.Token)
	// 	fmt.Println(User.Id, User.User_name, User.Email, User.Token)
	// }
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

var LoadRegister bool

var LoadLogin bool

func HandleTest(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
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

func HandleFilters(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
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
