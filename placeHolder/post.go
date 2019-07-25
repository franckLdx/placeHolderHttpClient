package placeHolder

type Post struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

type Posts []Post

func FindPostById(posts Posts, id int) *Post {
	predicat := func(post Post) bool { return post.Id == id }
	return find(posts, predicat)
}

func find(posts Posts, predicat func(Post) bool) *Post {
	for _, post := range posts {
		if predicat(post) {
			return &post
		}
	}
	return nil
}
