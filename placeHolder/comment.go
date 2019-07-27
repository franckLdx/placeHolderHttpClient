package placeHolder

type Comment struct {
	PostId int
	Id     int
	Name   string
	Email  string
	Body   string
}

type Comments []Comment
type CommentsMap map[int]Comments

type CommentsResult struct {
	PostId   int
	Comments Comments
	Err      error
}
