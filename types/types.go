package types

type Server struct {
	Url string
}

type Post struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

type Comment struct {
	OostId int
	Id     int
	Name   string
	Email  string
	Body   string
}
