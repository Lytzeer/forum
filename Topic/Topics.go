package forum

import (
	"database/sql"
	fd "forum/Datas"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var filedb string = "./database/database.db"

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
		err := rows.Scan(&topic.TopicID, &topic.TopicTitle, &topic.TopicMessage, &topic.TopicTime, &topic.TopicAuthor, &topic.TopicCategory)
		if err != nil {
			panic(err)
		}
		//fmt.Println(topic.TopicTitle, topic.TopicMessage, topic.TopicTime, topic.TopicAuthor, topic.TopicCategorie)
		topics = append(topics, topic)
	}
	defer rows.Close()
	return topics
}
