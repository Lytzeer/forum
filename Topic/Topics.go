package forum

import (
	"database/sql"
	"fmt"
	fd "forum/Datas"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var filedb string = "./database/forum.db"

func GetTopics() []fd.Topic {
	db, err := sql.Open("sqlite3", filedb)
	if err != nil {
		log.Fatal(err)
	}
	request_select := "SELECT * FROM Topics"
	rows, err := db.Query(request_select)
	if err != nil {
		log.Fatal(err)
	}
	var topics []fd.Topic
	for rows.Next() {
		var topic fd.Topic
		err := rows.Scan(&topic.TopicID, &topic.TopicTitle, &topic.TopicDate, &topic.TopicCreatorID)
		if err != nil {
			panic(err)
		}
		fmt.Println(topic.TopicID, topic.TopicTitle, topic.TopicCreator, topic.TopicDate, topic.TopicCreatorID)
		topics = append(topics, topic)
	}
	defer rows.Close()
	return topics
}
