package forum

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func AddComment(phrase string, user string, commentid int) string {
	if phrase == "" || user == "" {
		return "Empty field"
	} else {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		defer db.Close()
		stmt, err := db.Prepare("INSERT INTO Comments(title, like, dislike,topicid) VALUES (?,?,?,?)")
		CheckErr(err)
		defer stmt.Close()
		_, err = stmt.Exec(phrase, 0, 0, commentid)
		CheckErr(err)

		return "Comment added"
	}
}
