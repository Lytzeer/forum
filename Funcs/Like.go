package forum

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Like(message string) {
	if message == "" {
		return
	} else {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		request_like := "UPDATE Comments SET like = like + 1 WHERE message = '" + message + "'"
		_, err = db.Exec(request_like)
		CheckErr(err)
		db.Close()
	}
}
