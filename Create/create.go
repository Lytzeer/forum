package forum

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var filedb string = "./database/forum.db"

func Create(phrase string, user string) {

	db, err := sql.Open("sqlite3", filedb)
	if err != nil {
		log.Fatal(err)
	}
	request_message := "INSERT INTO (TopicMessage) VALUES ('" + phrase + "')"
	db.Exec(request_message)

	// request_count := "COUNT* id"
	// _, val := db.Exec(request_count)

	// request_id := "INSERT INTO (TopicID) VALUES ('" + val + "')"
	// db.Exec(request_id)

	request_username := "INSERT INTO (username) VALUES ('" + user + "')"
	db.Exec(request_username)

	cici := time.Now().Format("2006-01-02 15:04:05")
	request_date := "INSERT INTO (date) VALUES ('" + cici + "')"
	db.Exec(request_date)

}
