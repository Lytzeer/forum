package forum

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var filedb string = "./database/forum.db"

func Create(phrase string) {

	db, err := sql.Open("sqlite3", filedb)
	if err != nil {
		log.Fatal(err)
	}
	request_message := "INSERT INTO (TopicMessage) VALUES (+create+)"
	db.Exec(request_message)

	request_count := "COUNT* id"
	db.Exec(request_count)
	request_id := "INSERT INTO (TopicID) VALUES (+id+1+)"
	db.Exec(request_id)

	request_username := "INSERT INTO (username) VALUES (+User_name+)"
	db.Exec(request_username)

	cici := time.Now().Format("2006-01-02 15:04:05")
	request_caca := "INSERT INTO (date) VALUES ('" + cici + "')"
	request_caca.Exec

}
