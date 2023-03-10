package forum

type ProfileDatas struct {
	User          User
	Topics        []Topic
	LikedTopics   []Topic
	LikedComments []Comment
}

//Rajouter les posts likés et les posts commentés après avoir résolu les likes infinis
