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

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		ff.Error404(w, r)
		return
	} else {
		if LoadLogin == false {
			var tmpl *template.Template
			tmpl = template.Must(template.ParseFiles("./static/login.html"))
			err := tmpl.Execute(w, nil)
			ff.CheckErr(err)
			LoadLogin = true
			return
		} else {
			name, password := r.FormValue("mail-name"), r.FormValue("password-login")
			User.User_name = name
			var tmpl *template.Template
			Error.Error = ff.Login(name, password)
			if Error.Error == "Login successful" {
				var ProfileDatas fd.ProfileDatas
				var tmpl *template.Template
				ProfileDatas.LikedTopics, ProfileDatas.LikedComments, User.Token, ProfileDatas.Topics, ProfileDatas.User = ff.GetTopicsLiked(User.User_name), ff.GetCommentsLiked(User.User_name), ff.SetCookie(w, r, User.User_name), ff.GetTopics(), User
				tmpl = template.Must(template.ParseFiles("./static/profile.html"))
				LoadLogin = false
				err := tmpl.Execute(w, ProfileDatas)
				ff.CheckErr(err)
			} else {
				tmpl = template.Must(template.ParseFiles("./static/login.html"))
				err := tmpl.Execute(w, Error)
				ff.CheckErr(err)
				Error.Error = ""
			}
			return
		}
	}
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		ff.Error404(w, r)
		return
	} else {
		if LoadRegister == false {
			LoadRegister = true
			var tmpl *template.Template
			tmpl = template.Must(template.ParseFiles("./static/register.html"))
			err := tmpl.Execute(w, nil)
			ff.CheckErr(err)
			return
		} else {
			var tmpl *template.Template
			tmpl = template.Must(template.ParseFiles("./static/register.html"))
			name, mail, password, confipass := r.FormValue("name"), r.FormValue("mail"), r.FormValue("password-register"), r.FormValue("conf-password-register")
			Error.Error = ff.Register(name, mail, password, confipass)
			if Error.Error == "Register successful" {
				LoadRegister = false
				tmpl = template.Must(template.ParseFiles("./static/login.html"))
				err := tmpl.Execute(w, nil)
				ff.CheckErr(err)
				return
			} else {
				err := tmpl.Execute(w, Error)
				ff.CheckErr(err)
				return
			}
		}
	}
}

func HandleProfile(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
	if r.URL.Path != "/profile" {
		ff.Error404(w, r)
		return
	} else {
		if User.Token == "" {
			http.Redirect(w, r, "/login", 302)
		} else {
			var tmpl *template.Template
			var ProfileDatas fd.ProfileDatas
			ProfileDatas.LikedTopics, ProfileDatas.LikedComments, ProfileDatas.Topics, ProfileDatas.User = ff.GetTopicsLiked(User.User_name), ff.GetCommentsLiked(User.User_name), ff.GetTopics(), User
			tmpl = template.Must(template.ParseFiles("./static/profile.html"))
			err := tmpl.Execute(w, ProfileDatas)
			ff.CheckErr(err)
			return
		}
	}
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
	if r.URL.Path != "/logout" {
		ff.Error404(w, r)
		return
	} else {
		User.User_name, User.Token, User.SignIn, User.Id = "", "", false, 0
		T.User = User
		ff.DeleteCookie(w, r)
		http.Redirect(w, r, "/", 302)
		return
	}
}

func HandleEditProfile(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
	if r.URL.Path != "/editprofile" {
		ff.Error404(w, r)
		return
	} else {
		if User.Token == "" {
			http.Redirect(w, r, "/login", 302)
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
			var ProfileDatas fd.ProfileDatas
			var tmpl *template.Template
			ProfileDatas.Topics, ProfileDatas.User, ProfileDatas.LikedTopics, ProfileDatas.LikedComments = ff.GetTopics(), User, ff.GetTopicsLiked(User.User_name), ff.GetCommentsLiked(User.User_name)
			tmpl = template.Must(template.ParseFiles("./static/editprofile.html"))
			err := tmpl.Execute(w, ProfileDatas)
			ff.CheckErr(err)
			return
		}
	}
}

func HandleNotif(w http.ResponseWriter, r *http.Request) {
	LoadLogin, LoadRegister = false, false
	if r.URL.Path != "/notif" {
		ff.Error404(w, r)
		return
	} else {
		if User.Token == "" {
			http.Redirect(w, r, "/login", 302)
		} else {
			Notifications := ff.GetNotifs(User.User_name)
			var tmpl *template.Template
			tmpl = template.Must(template.ParseFiles("./static/notif.html"))
			err := tmpl.Execute(w, Notifications)
			ff.CheckErr(err)
			return
		}
	}
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

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}

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
