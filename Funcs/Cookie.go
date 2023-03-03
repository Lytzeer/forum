package forum

import (
	"database/sql"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"

	uuid "github.com/satori/go.uuid"
)

func SetCookie(w http.ResponseWriter, r *http.Request, username string) string {
	session_token := uuid.NewV4()
	cookie := &http.Cookie{
		Name:   "session",
		Value:  session_token.String(),
		MaxAge: 86400,
	}
	http.SetCookie(w, cookie)

	db, err := sql.Open("sqlite3", filedb)
	CheckErr(err)

	request_token, err := db.Prepare("UPDATE User SET token = ? WHERE username = ?")
	CheckErr(err)

	request_token.Exec(session_token.String(), username)
	defer db.Close()
	return session_token.String()
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:    "session",
		Value:   "",
		MaxAge:  -1,
		Expires: time.Now().Add(-1 * time.Hour),
	}
	http.SetCookie(w, cookie)
}

func GetCookie(r *http.Request) string {
	cookie, err := r.Cookie("session")
	if err != nil {
		return ""
	}
	return cookie.Value
}
