package forum

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var filedb string = "./database/forum.db"

func Login(username string, password string) int {
	if username == "" || password == "" {
		return 1
	} else {
		db, err := sql.Open("sqlite3", filedb)
		if err != nil {
			log.Fatal(err)
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(hash))
		request_select := "SELECT username, email, password FROM User WHERE username = '" + username + "' OR email = '" + username + "' AND password = '" + string(hash) + "'"
		rows, err := db.Query(request_select)
		if err != nil {
			log.Fatal(err)
		}
		var username_db string
		var email_db string
		var password_db string

		for rows.Next() {
			err = rows.Scan(&username_db, &email_db, &password_db)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(username_db, email_db, password_db)
		}
		if username_db == "" || email_db == "" || password_db == "" {
			return 2
		}
		rows.Close()
		return 0
	}

}
