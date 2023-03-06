package forum

import (
	"database/sql"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func ModifyComment(phrase string, user string, commentID int) string {
	if user == "" || commentID == 0 || phrase == "" {
		return "missing field"
	} else {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		request_delete_topic := ("SELECT creatorname FROM Comments WHERE id='" + strconv.Itoa(commentID) + "'")
		rows, err := db.Query(request_delete_topic)
		CheckErr(err)
		var usr string
		for rows.Next() {
			err = rows.Scan(&usr)
			CheckErr(err)
		}
		if usr != user {
			return "je mange mon caca"
		} else {
			request_delete_comments, err := db.Prepare("UPDATE Comments SET title=? WHERE id=?")
			CheckErr(err)
			request_delete_comments.Exec(phrase, commentID)
			db.Close()
			return "gg bro"
		}
	}
}
