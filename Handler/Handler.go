package forum

import (
	"encoding/json"
	"fmt"
	fd "forum/Datas"
	ff "forum/Funcs"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

type displayerror struct {
	Leprobleme string
}

var t displayerror
var User fd.User
var Topics []fd.Topic
var Topic fd.Topic

var modifycommentid int
var modifytopicid int

var T fd.Topics
var T2 fd.TopicInfos

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if User.SignIn == false {
		token, _ := r.Cookie("session")
		User.Id, User.User_name, User.Email, User.Token = ff.CheckToken(token.Value)
		fmt.Println(token.Value == User.Token)
		fmt.Println(User.Token)
		fmt.Println(User.Id, User.User_name, User.Email, User.Token)
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
			title, message, categorie := r.FormValue("title"), r.FormValue("message"), r.FormValue("categories")
			fmt.Println(title, message, categorie, User.User_name)
			t.Leprobleme = ff.Create(message, User.User_name, title, categorie)
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
			commentId, message := r.FormValue("modify"), r.FormValue("message")
			if modifycommentid == 0 {
				modifycommentid, _ = strconv.Atoi(commentId)
			}
			commentIdstr, _ := strconv.Atoi(commentId)
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
			comment, title, topicid := r.FormValue("message"), r.FormValue("title"), r.FormValue("modify")
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
			ff.Like(idint)
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

var (
	oauthConf = &oauth2.Config{
		ClientID:     "1210097736303314",
		ClientSecret: "386a018e3f1fd63e3d70a5a0a65fcc65",
		RedirectURL:  "http://localhost:8080/oauth2callback",
		Scopes:       []string{"public_profile"},
		Endpoint:     facebook.Endpoint,
	}
	oauthStateString = "thisshouldberandom"
)

func HandleLoginFacebook(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/loginfacebook" {
		ff.Error404(w, r)
		return
	} else {
		Url, err := url.Parse(oauthConf.Endpoint.AuthURL)
		ff.CheckErr(err)
		parameters := url.Values{}
		parameters.Add("client_id", oauthConf.ClientID)
		parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
		parameters.Add("redirect_uri", oauthConf.RedirectURL)
		parameters.Add("response_type", "code")
		parameters.Add("state", oauthStateString)
		Url.RawQuery = parameters.Encode()
		url := Url.String()
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func HandleFacebookCallback(w http.ResponseWriter, r *http.Request) {

	code := r.FormValue("code")

	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	ff.CheckErr(err)

	resp, err := http.Get("https://graph.facebook.com/me?access_token=" +
		url.QueryEscape(token.AccessToken))
	ff.CheckErr(err)
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	ff.CheckErr(err)

	var fbUser fd.Facebook
	err = json.Unmarshal(response, &fbUser)
	ff.CheckErr(err)

	User.User_name = fbUser.Name
	idint, _ := strconv.Atoi(fbUser.ID)
	User.Id = idint

	User.Token = ff.SetCookie(w, r, User.User_name)

	http.Redirect(w, r, "/profile", http.StatusTemporaryRedirect)
}

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}

const clientID = "5fe1b31b4f9577bc2463"
const clientSecret = "f19d67ab63b0f452c5a9109f445fc99b7cfe098e"

func HandleLoginGithub(w http.ResponseWriter, r *http.Request) {
	httpClient := http.Client{}
	if r.URL.Path != "/oauth/redirect" {
		ff.Error404(w, r)
		return
	} else {
		err := r.ParseForm()
		ff.CheckErr(err)
		code := r.FormValue("code")

		// Next, lets for the HTTP request to call the github oauth endpoint
		// to get our access token
		reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
		req, err := http.NewRequest(http.MethodPost, reqURL, nil)
		ff.CheckErr(err)
		req.Header.Set("accept", "application/json")

		// Send out the HTTP request
		res, err := httpClient.Do(req)
		ff.CheckErr(err)
		defer res.Body.Close()

		var t OAuthAccessResponse

		err = json.NewDecoder(res.Body).Decode(&t)
		ff.CheckErr(err)

		w.Header().Set("Location", "/profile?access_token="+t.AccessToken)
		w.WriteHeader(http.StatusFound)

		req, err = http.NewRequest("GET", "https://api.github.com/user", nil)
		req.Header.Set("Accept", "application/vnd.github.v3+json")
		req.Header.Set("Authorization", "token "+t.AccessToken)

		response, err := httpClient.Do(req)
		ff.CheckErr(err)
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		ff.CheckErr(err)

		var tttt fd.Github

		err = json.Unmarshal(body, &tttt)
		ff.CheckErr(err)

		User.User_name = tttt.Login
		User.Id = tttt.Id

		User.Token = ff.SetCookie(w, r, User.User_name)

	}
}
