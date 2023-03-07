package forum

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, password string) string {
	if username == "" || password == "" {
		return "Please fill all the fields"
	} else {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		request_select := "SELECT username, email, password,token FROM User WHERE username = '" + username + "' OR email = '" + username + "'"
		rows, err := db.Query(request_select)
		CheckErr(err)
		var username_db string
		var email_db string
		var password_db string
		var token_db string
		for rows.Next() {
			err = rows.Scan(&username_db, &email_db, &password_db, &token_db)
			CheckErr(err)
		}
		if bcrypt.CompareHashAndPassword([]byte(password_db), []byte(password)) != nil {
			return "Wrong username or password"
		}
		if username_db == "" || email_db == "" || password_db == "" {
			return "Wrong username or password"
		}
		rows.Close()
		return "Login successful"
	}

}
