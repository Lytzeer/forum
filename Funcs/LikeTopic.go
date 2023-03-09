package forum

import (
	"database/sql"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func LikeTopic(id int) {
	if id == 0 {
		return
	} else {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		idstr := strconv.Itoa(id)
		request_like := "UPDATE Topics SET like = like + 1 WHERE id = '" + idstr + "'"
		_, err = db.Exec(request_like)
		CheckErr(err)
		db.Close()
	}
}

func DislikeTopic(id int) {
	if id == 0 {
		return
	} else {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		request_dislike := "UPDATE Topics SET dislike = dislike + 1 WHERE id = '" + strconv.Itoa(id) + "'"
		_, err = db.Exec(request_dislike)
		CheckErr(err)
		db.Close()
	}
}
