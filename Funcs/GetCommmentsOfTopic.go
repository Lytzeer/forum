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
		request_comments := "SELECT id,title,like,dislike,topicid,creatorname FROM comments WHERE topicid=" + idstr + ";"
		rows, err := db.Query(request_comments)
		CheckErr(err)
		for rows.Next() {
			var comment fd.Comment
			err = rows.Scan(&comment.Id, &comment.Title, &comment.Likes, &comment.Dislikes, &comment.TopicId, &comment.CreatorName)
			CheckErr(err)
			comments = append(comments, comment)
		}
		return comments
	}
}
