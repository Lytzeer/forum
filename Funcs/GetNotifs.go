package forum

import (
	"database/sql"

	fd "forum/Datas"
)

func GetNotifs(user string) []fd.Notif {
	if user == "" {
		return nil
	} else {
		var Notifs []fd.Notif
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		request_comments := "SELECT id, date, user, str FROM Notif WHERE user = '" + user + "'"
		rows, err := db.Query(request_comments)
		CheckErr(err)
		for rows.Next() {
			var notif fd.Notif
			err = rows.Scan(&notif.NotifID, &notif.NotifDate, &notif.NotifDest, &notif.NotifMessage)
			CheckErr(err)
			Notifs = append(Notifs, notif)
		}
		return Notifs
	}
}
