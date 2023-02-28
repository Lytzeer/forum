package forum

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var filedb string = "./database/database.db"

func Create(phrase string, user string, title string, categorie string) string {
	if phrase == "" || user == "" || title == "" || categorie == "" {
		return "Empty fields"
	} else {
		db, err := sql.Open("sqlite3", filedb)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		//reqeuete pour ajouter un topic
		stmt, err := db.Prepare("INSERT INTO Topics (title, message, date, categorie, creatorname) VALUES (?, ?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(title, phrase, time.Now().String(), categorie, user)
		if err != nil {
			log.Fatal(err)
		}

		return "Topic created"
	}
}
