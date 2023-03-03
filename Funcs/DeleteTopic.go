package forum

import (
	"database/sql"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func DeleteTopic(user string, topicID int) string {
	if user == "" || topicID == 0 {
		return "missing username"
	} else {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		request_delete_topic := ("SELECT creatorname FROM Topics WHERE id='" + strconv.Itoa(topicID) + "'")
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
			request_delete_comments, err := db.Prepare("DELETE FROM Topics WHERE id=?")
			CheckErr(err)
			request_delete_comments.Exec(topicID)
			db.Close()
			request_delete_comments, err = db.Prepare("DELETE FROM Comments WHERE topicid=?")
			CheckErr(err)
			request_delete_comments.Exec(topicID)
			db.Close()
			return "gg bro"
		}
	}
}
