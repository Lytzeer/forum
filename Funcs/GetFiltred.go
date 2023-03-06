package forum

import (
	"database/sql"

	fd "forum/Datas"
)

func GetFiltred(categorie string) []fd.Topic {
	if categorie == "" {
		return nil
	} else {
		var Topics []fd.Topic
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		request_filtre := "SELECT id,title,message,  date, categorie, creatorname FROM Topics WHERE categorie = '" + categorie + "'"
		rows, err := db.Query(request_filtre)
		CheckErr(err)
		for rows.Next() {
			var topic fd.Topic
			err := rows.Scan(&topic.TopicID, &topic.TopicTitle, &topic.TopicMessage, &topic.TopicTime, &topic.TopicCategory, &topic.TopicAuthor)
			CheckErr(err)
			Topics = append(Topics, topic)
		}
		return Topics
	}
}
