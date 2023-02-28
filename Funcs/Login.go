package forum

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, password string) string {
	if username == "" || password == "" {
		return "Please fill all the fields"
	} else {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(hash))
		request_select := "SELECT username, email, password FROM User WHERE username = '" + username + "' OR email = '" + username + "' AND password = '" + string(hash) + "'"
		rows, err := db.Query(request_select)
		CheckErr(err)
		var username_db string
		var email_db string
		var password_db string

		for rows.Next() {
			err = rows.Scan(&username_db, &email_db, &password_db)
			CheckErr(err)
			fmt.Println(username_db, email_db, password_db)
		}
		if username_db == "" || email_db == "" || password_db == "" {
			return "Wrong username or password"
		}
		rows.Close()
		return "Login successful"
	}

}
