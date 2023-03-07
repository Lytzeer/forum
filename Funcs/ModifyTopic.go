package forum

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func ModifyTopic(title string, phrase string, user string, topicID int) string {
	if user == "" || topicID == 0 || phrase == "" || title == "" {
		return "missing field"
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
		fmt.Println(usr, user)
		if usr != user {
			return "je mange mon caca" + usr + user
		} else {
			request_delete_comments, err := db.Prepare("UPDATE Topics SET title=?, message=? WHERE id=?")
			CheckErr(err)
			request_delete_comments.Exec(title, phrase, topicID)
			db.Close()
			return "gg bro"
		}
	}
}
