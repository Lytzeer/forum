package forum

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func LikeTopic(id int, user string, topicauthor string) {
	if id == 0 || user == "" {
		return
	} else {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		request_get_liked_topic := "SELECT likedTopics FROM User WHERE username='" + user + "'"
		CheckErr(err)
		rows, err := db.Query(request_get_liked_topic)
		CheckErr(err)
		var likedTopics string
		for rows.Next() {
			err = rows.Scan(&likedTopics)
			CheckErr(err)
		}
		idstr := strconv.Itoa(id)
		if strings.Contains(likedTopics, idstr) {
			request_like := "UPDATE Topics SET like = like - 1 WHERE id = '" + idstr + "'"
			DeleteLikeUser(user, idstr, "likedTopics")
			_, err = db.Exec(request_like)
		} else {
			request_like := "UPDATE Topics SET like = like + 1 WHERE id = '" + idstr + "'"
			AddLikeUser(user, idstr, "likedTopics")
			_, err = db.Exec(request_like)
			notif, err := db.Prepare("INSERT INTO Notif(date, user, str) VALUES (?,?,?)")
			CheckErr(err)
			defer notif.Close()
			_, err = notif.Exec(time.Now().String(), topicauthor, user+" a commenté votre post")
			CheckErr(err)
			defer db.Close()
		}
	}
}

func DislikeTopic(id int, user string, topicauthor string) {
	if id == 0 || user == "" {
		return
	} else {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		request_get_disliked_topic := "SELECT dislikedTopics FROM User WHERE username='" + user + "'"
		CheckErr(err)
		rows, err := db.Query(request_get_disliked_topic)
		CheckErr(err)
		var dislikedTopics string
		for rows.Next() {
			err = rows.Scan(&dislikedTopics)
			CheckErr(err)
		}
		idstr := strconv.Itoa(id)
		if strings.Contains(dislikedTopics, idstr) {
			request_dislike := "UPDATE Topics SET dislike = dislike - 1 WHERE id = '" + idstr + "'"
			DeleteLikeUser(user, idstr, "dislikedTopics")
			_, err = db.Exec(request_dislike)
		} else {
			request_dislike := "UPDATE Topics SET dislike = dislike + 1 WHERE id = '" + idstr + "'"
			AddLikeUser(user, idstr, "dislikedTopics")
			_, err = db.Exec(request_dislike)
			notif, err := db.Prepare("INSERT INTO Notif(date, user, str) VALUES (?,?,?)")
			CheckErr(err)
			defer notif.Close()
			_, err = notif.Exec(time.Now().String(), topicauthor, user+" a commenté votre post")
			CheckErr(err)
			defer db.Close()
		}
	}
}
