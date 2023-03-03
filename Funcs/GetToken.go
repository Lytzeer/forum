package forum

import (
	"database/sql"
	"fmt"
)

func CheckToken(token string) (int, string, string, string) {
	db, err := sql.Open("sqlite3", filedb)
	CheckErr(err)
	request_count := "SELECT COUNT(*) FROM User WHERE token = ?"
	rows, err := db.Query(request_count, token)
	CheckErr(err)
	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		CheckErr(err)
	}
	if count == 1 {
		request_getinfos, err := db.Prepare("SELECT id, username, email FROM User WHERE token = ?")
		CheckErr(err)
		rows, err := request_getinfos.Query(token)
		var id int
		var username string
		var email string
		CheckErr(err)
		for rows.Next() {
			err = rows.Scan(&id, &username, &email)
			CheckErr(err)
		}
		fmt.Println(id, username, email)
		return id, username, email, token
	}
	return 0, "", "", ""
}
