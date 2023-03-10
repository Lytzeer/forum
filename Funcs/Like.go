package forum

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Like(id int, user string, comment_author string) {
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
		db, err = sql.Open("sqlite3", filedb)
		idstr := strconv.Itoa(id)
		CheckErr(err)
		if strings.Contains(likedTopics, idstr) {
			request_like := "UPDATE Comments SET like = like - 1 WHERE id = '" + idstr + "'"
			DeleteLikeUser(user, idstr, "likedTopics")
			_, err = db.Exec(request_like)
		} else {
			request_like := "UPDATE Comments SET like = like + 1 WHERE id = '" + idstr + "'"
			AddLikeUser(user, idstr, "likedTopics")
			_, err = db.Exec(request_like)
			notif, err := db.Prepare("INSERT INTO Notif(date, user, str) VALUES (?,?,?)")
			CheckErr(err)
			defer notif.Close()
			_, err = notif.Exec(time.Now().String(), comment_author , user+" a liké votre post")
			CheckErr(err)
			defer db.Close()
		}
		CheckErr(err)
		db.Close()
	}
}

func Dislike(id int, user string, comment_author string) {
	if id == 0 {
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
		db, err = sql.Open("sqlite3", filedb)
		idstr := strconv.Itoa(id)
		CheckErr(err)
		if strings.Contains(dislikedTopics, idstr) {
			request_dislike := "UPDATE Comments SET dislike = dislike - 1 WHERE id = '" + idstr + "'"
			DeleteLikeUser(user, idstr, "dislikedTopics")
			_, err = db.Exec(request_dislike)
		} else {
			request_dislike := "UPDATE Comments SET dislike = dislike + 1 WHERE id = '" + idstr + "'"
			AddLikeUser(user, idstr, "dislikedTopics")
			_, err = db.Exec(request_dislike)
			notif, err := db.Prepare("INSERT INTO Notif(date, user, str) VALUES (?,?,?)")
			CheckErr(err)
			defer notif.Close()
			_, err = notif.Exec(time.Now().String(), comment_author , user+" a liké votre post")
			CheckErr(err)
			defer db.Close()
		}
	}
}

func AddLikeUser(user string, id string, typelike string) string {
	db, err := sql.Open("sqlite3", filedb)
	CheckErr(err)
	request_get_liked_topic := "SELECT " + typelike + " FROM User WHERE username='" + user + "'"
	CheckErr(err)
	rows, err := db.Query(request_get_liked_topic)
	CheckErr(err)
	var likedTopics string
	for rows.Next() {
		err = rows.Scan(&likedTopics)
		CheckErr(err)
	}
	listedLikedTopics := strings.Split(likedTopics, "-")
	strings.TrimSpace(likedTopics)
	fmt.Println(listedLikedTopics)
	var newLikedTopics string
	for _, ch := range listedLikedTopics {
		if string(ch) != id && string(ch) != "" {
			newLikedTopics += string(ch)
			newLikedTopics += "-"
		}
	}
	newLikedTopics += id
	request_update_liked_topic := "UPDATE User SET " + typelike + " = '" + newLikedTopics + "' WHERE username='" + user + "'"
	_, err = db.Exec(request_update_liked_topic)
	CheckErr(err)
	return newLikedTopics
}

func DeleteLikeUser(user string, id string, typelike string) string {
	db, err := sql.Open("sqlite3", filedb)
	CheckErr(err)
	request_get_liked_topic := "SELECT " + typelike + " FROM User WHERE username='" + user + "'"
	CheckErr(err)
	rows, err := db.Query(request_get_liked_topic)
	CheckErr(err)
	var likedTopics string
	for rows.Next() {
		err = rows.Scan(&likedTopics)
		CheckErr(err)
	}
	listedLikedTopics := strings.Split(likedTopics, "-")
	strings.TrimSpace(likedTopics)
	var newLikedTopics string
	for _, ch := range listedLikedTopics {
		if string(ch) != id && string(ch) != "" {
			newLikedTopics += string(ch)
			newLikedTopics += "-"
		}
	}
	request_update_liked_topic := "UPDATE User SET " + typelike + " = '" + newLikedTopics + "' WHERE username='" + user + "'"
	_, err = db.Exec(request_update_liked_topic)
	CheckErr(err)
	return newLikedTopics
}
