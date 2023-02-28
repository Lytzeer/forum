package forum

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var filedb string = "./database/database.db"

func Register(name string, mail string, password string, confpass string) string {
	if name == "" || mail == "" || password == "" || confpass == "" {
		return "Please fill all the fields"
	} else {
		db, err := sql.Open("sqlite3", filedb)
		if err != nil {
			fmt.Println(err)
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
		}
		request_count := "SELECT COUNT(*) FROM User WHERE username = '" + name + "' OR email = '" + mail + "' AND password = '" + string(hash) + "'"
		rows, err := db.Query(request_count)

		if err != nil {
			log.Fatal(err)
		}
		var count int

		for rows.Next() {
			err = rows.Scan(&count)
			if err != nil {
				log.Fatal(err)
			}
		}
		rows.Close()

		if count == 0 {
			if err := bcrypt.CompareHashAndPassword(hash, []byte(confpass)); err == nil {
				request_register, err := db.Prepare("INSERT INTO User (username,email, password) VALUES ('" + name + "', '" + mail + "', '" + string(hash) + "')")
				if err != nil {
					log.Fatal(err)
				}
				request_register.Exec()
			} else {
				return "Passwords don't match"
			}
		}
		return "Register successful"

	}

}
