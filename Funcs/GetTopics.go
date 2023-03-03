package forum

import (
	"database/sql"
	fd "forum/Datas"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var filedb string = "./database/database.db"

func GetTopics() []fd.Topic {
	db, err := sql.Open("sqlite3", filedb)
	CheckErr(err)
	request_select := "SELECT * FROM Topics"
	rows, err := db.Query(request_select)
	CheckErr(err)
	var topics []fd.Topic
	for rows.Next() {
		var topic fd.Topic
		err := rows.Scan(&topic.TopicID, &topic.TopicTitle, &topic.TopicMessage, &topic.TopicTime, &topic.TopicCategory, &topic.TopicAuthor)
		CheckErr(err)
		topics = append(topics, topic)
	}
	defer rows.Close()
	return topics
}

func GetOneTopics(id int) fd.Topic {
	var topic fd.Topic
	db, err := sql.Open("sqlite3", filedb)
	CheckErr(err)
	request_select := "SELECT * FROM Topics WHERE id = " + strconv.Itoa(id)
	rows, err := db.Query(request_select)
	CheckErr(err)
	for rows.Next() {
		err := rows.Scan(&topic.TopicID, &topic.TopicTitle, &topic.TopicMessage, &topic.TopicTime, &topic.TopicCategory, &topic.TopicAuthor)
		CheckErr(err)
	}
	defer rows.Close()
	return topic
}
