package forum

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Dislike(message string) {
	if message == "" {
		return
	} else {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		request_dislike := "UPDATE Comments SET dislike = dislike + 1 WHERE message = '" + message + "'"
		_, err = db.Exec(request_dislike)
		CheckErr(err)
		db.Close()
	}
}
