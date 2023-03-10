package forum

import (
	"database/sql"
	"fmt"
	fd "forum/Datas"
	"strconv"
	"strings"

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
		err := rows.Scan(&topic.TopicID, &topic.TopicTitle, &topic.TopicMessage, &topic.TopicTime, &topic.TopicCategory, &topic.TopicAuthor, &topic.Likes, &topic.Dislikes)
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
		err := rows.Scan(&topic.TopicID, &topic.TopicTitle, &topic.TopicMessage, &topic.TopicTime, &topic.TopicCategory, &topic.TopicAuthor, &topic.Likes, &topic.Dislikes)
		CheckErr(err)
	}
	defer rows.Close()
	return topic
}

func GetTopicsLiked(user string) []fd.Topic {
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

	fmt.Println(likedTopics)

	liststr := strings.Split(likedTopics, "-")
	var topics []fd.Topic
	for _, str := range liststr {
		var topic fd.Topic
		idint, _ := strconv.Atoi(str)
		topic = GetOneTopics(idint)
		fmt.Println(topic.Comments)
		topics = append(topics, topic)
	}
	defer rows.Close()
	return topics
}
