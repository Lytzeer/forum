package forum

import (
	"database/sql"
	"strconv"

	fd "forum/Datas"
)

func GetCommmentsOfTopic(topicId int) []fd.Comment {
	if topicId == 0 {
		return nil
	} else {
		var comments []fd.Comment
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		var idstr string = strconv.Itoa(topicId)
		request_comments := "SELECT * FROM comments WHERE topic_id=" + idstr + ";"
		rows, err := db.Query(request_comments)
		CheckErr(err)
		for rows.Next() {
			var comment fd.Comment
			err = rows.Scan(&comment.Id, &comment.Title, &comment.Like, &comment.Dislike, &topicId)
			CheckErr(err)
			comments = append(comments, comment)
		}
		return comments
	}
}
