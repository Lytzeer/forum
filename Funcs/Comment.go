package forum

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	fd "forum/Datas"

	_ "github.com/mattn/go-sqlite3"
)

func AddComment(phrase string, user string, commentid int, topicauthor string) string {
	if phrase == "" || user == "" {
		return "Empty field"
	} else {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
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
		defer db.Close()

		return "Comment added"
	}
}

func GetOneComment(id int) fd.Comment {
	db, err := sql.Open("sqlite3", filedb)
	CheckErr(err)
	request_get_one_topic := ("SELECT * FROM Comments WHERE id='" + strconv.Itoa(id) + "'")
	rows, err := db.Query(request_get_one_topic)
	CheckErr(err)
	var comment fd.Comment
	for rows.Next() {
		err = rows.Scan(&comment.Id, &comment.Title, &comment.Likes, &comment.Dislikes, &comment.TopicId, &comment.CreatorName)
		CheckErr(err)
	}
	defer db.Close()
	return comment
}

func GetCommentsLiked(user string) []fd.Comment {
	db, err := sql.Open("sqlite3", filedb)
	CheckErr(err)

	request_get_liked_topic := "SELECT likedComments FROM User WHERE username='" + user + "'"
	CheckErr(err)
	rows, err := db.Query(request_get_liked_topic)
	CheckErr(err)
	var likedComments string
	for rows.Next() {
		err = rows.Scan(&likedComments)
		CheckErr(err)
	}
	defer rows.Close()

	liststr := strings.Split(likedComments, "-")
	var comments []fd.Comment
	for _, str := range liststr {
		if str != "" && str != " " {
			i, err := strconv.Atoi(str)
			CheckErr(err)
			request_select_topic := "SELECT * FROM Comments WHERE id = " + strconv.Itoa(i)
			rows, err := db.Query(request_select_topic)
			CheckErr(err)
			for rows.Next() {
				var comment fd.Comment
				err := rows.Scan(&comment.Id, &comment.Title, &comment.Likes, &comment.Dislikes, &comment.TopicId, &comment.CreatorName)
				CheckErr(err)
				comments = append(comments, comment)
			}
			defer rows.Close()
		}
	}
	return comments
}
