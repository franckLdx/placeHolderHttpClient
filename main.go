package main

import (
	"log"
	"placeHolderHttpClient/placeHolder"
	placeholder "placeHolderHttpClient/placeHolder"
	"placeHolderHttpClient/tools"
)

func main() {
	server := getPlaceHolder()

	posts := getPosts(server)

	commentsChan := make(chan *placeHolder.CommentsResult)
	defer close(commentsChan)
	go getComments(server, posts, commentsChan)

	print(posts, commentsChan)

	log.Println("Bye Bye")
}

func getPlaceHolder() *placeHolder.Server {
	config := tools.LoadConfig()
	return &placeHolder.Server{Url: config.Server.Url}
}

func print(posts *[]placeHolder.Post, commentsChan <-chan *placeHolder.CommentsResult) {
	for count := 0; count < len(*posts); count++ {
		commentResult := <-commentsChan
		if commentResult.Err != nil {
			log.Fatal(commentResult.Err)
		}
		log.Println(commentResult.Comments)
		log.Println(count)
	}
}

func getPosts(server *placeHolder.Server) *[]placeHolder.Post {
	posts, err := server.GetPosts()
	if err != nil {
		log.Fatal(err)
	}
	return posts
}

func getComments(server *placeholder.Server, posts *[]placeholder.Post, commentsResult chan<- (*placeHolder.CommentsResult)) {
	for _, post := range *posts {
		go server.GoGetComments(post.Id, commentsResult)
	}
}
