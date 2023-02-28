package forum

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/mattn/go-sqlite3"
)

func Register(name string, mail string, password string, confpass string) string {
	if name == "" || mail == "" || password == "" || confpass == "" {
		return "Please fill all the fields"
	} else {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		CheckErr(err)
		request_count := "SELECT COUNT(*) FROM User WHERE username = '" + name + "' OR email = '" + mail + "' AND password = '" + string(hash) + "'"
		rows, err := db.Query(request_count)

		CheckErr(err)
		var count int

		for rows.Next() {
			err = rows.Scan(&count)
			CheckErr(err)
		}
		rows.Close()

		if count == 0 {
			if err := bcrypt.CompareHashAndPassword(hash, []byte(confpass)); err == nil {
				request_register, err := db.Prepare("INSERT INTO User (username,email, password) VALUES ('" + name + "', '" + mail + "', '" + string(hash) + "')")
				CheckErr(err)
				request_register.Exec()
			} else {
				return "Passwords don't match"
			}
		}
		return "Register successful"

	}

}
