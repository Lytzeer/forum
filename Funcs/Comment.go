package forum

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func AddComment(phrase string, user string, commentid int, topicauthor string) string {
	if phrase == "" || user == "" {
		return "Empty field"
	} else {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		defer db.Close()
		stmt, err := db.Prepare("INSERT INTO Comments(title, like, dislike,topicid, creatorname) VALUES (?,?,?,?,?)")
		CheckErr(err)
		defer stmt.Close()
		_, err = stmt.Exec(phrase, 0, 0, commentid, user)
		CheckErr(err)
		notif, err := db.Prepare("INSERT INTO Notif(date, user, str) VALUES (?,?,?)")
		CheckErr(err)
		defer notif.Close()
		_, err = notif.Exec(time.Now().String(), topicauthor, user+" à commenté votre post")
		CheckErr(err)

		return "Comment added"
	}
}
