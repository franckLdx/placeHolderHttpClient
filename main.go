package main

import (
	"log"
	"placeHolderHttpClient/placeHolder"
	placeholder "placeHolderHttpClient/placeHolder"
	"placeHolderHttpClient/tools"
)

func main() {
	placeHolder := getPlaceHolder()
	posts, err := placeHolder.GetPosts()
	if err != nil {
		log.Fatal(err)
	}
	printPosts(placeHolder, posts)
	log.Println("Bye Bye")
}

func getPlaceHolder() *placeHolder.Server {
	config := tools.LoadConfig()
	return &placeHolder.Server{Url: config.Server.Url}
}

func printPosts(server *placeHolder.Server, posts *[]placeHolder.Post) {
	commentsChan := make(chan *placeHolder.CommentsResult)
	defer close(commentsChan)
	go getComments(server, posts, commentsChan)

	for count := 0; count < len(*posts); count++ {
		commentResult := <-commentsChan
		if commentResult.Err != nil {
			log.Fatal(commentResult.Err)
		}
		log.Println(commentResult.Comments)
		log.Println(count)
	}
}

func getComments(server *placeholder.Server, posts *[]placeholder.Post, result chan<- (*placeHolder.CommentsResult)) {
	for _, post := range *posts {
		go server.GoGetComments(post.Id, result)
	}
}
