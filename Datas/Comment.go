package forum

type Comment struct {
	Id          int
	Title       string
	Likes       int
	Dislikes    int
	TopicId     int
	CreatorName string
}
