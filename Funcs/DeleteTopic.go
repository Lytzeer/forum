package forum

import "database/sql"

func DeleteTopic(user string, topicID int) {
	if user == "" || topicID == 0 {
		return
	} else {
		db, err := sql.Open("mysql", filedb)
		CheckErr(err)
		request_delete_topic, err := db.Prepare("DELETE FROM Topics WHERE id=? AND creatorname=?")
		CheckErr(err)
		request_delete_topic.Exec(topicID, user)
		request_delete_comments, err := db.Prepare("DELETE FROM Comments WHERE topicid=?")
		CheckErr(err)
		request_delete_comments.Exec(topicID)
		db.Close()
	}
}
