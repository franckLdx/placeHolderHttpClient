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

	printPostsWithComments(posts, commentsChan)

	log.Println("Bye Bye")
}

func getPlaceHolder() *placeHolder.Server {
	config := tools.LoadConfig()
	return &placeHolder.Server{Url: config.Server.Url}
}

func getPosts(server *placeHolder.Server) *placeHolder.Posts {
	posts, err := server.GetPosts()
	if err != nil {
		log.Fatal(err)
	}
	return posts
}

func getComments(server *placeholder.Server, posts *placeholder.Posts, commentsResult chan<- (*placeHolder.CommentsResult)) {
	for _, post := range *posts {
		go server.GoGetComments(post.Id, commentsResult)
	}
}

func printPostsWithComments(posts *placeHolder.Posts, commentsChan <-chan *placeHolder.CommentsResult) {
	for count := 0; count < len(*posts); count++ {
		commentResult := <-commentsChan
		if commentResult.Err != nil {
			log.Fatal(commentResult.Err)
		}
		post := placeHolder.FindPostById(*posts, commentResult.PostId)
		if post == nil {
			log.Fatal("Comments with no posts", commentResult.PostId)
		}
		printPostWithComments(post, &commentResult.Comments)
	}
}

func printPostWithComments(post *placeHolder.Post, comments *placeHolder.Comments) {
	log.Println("----------------------------------------------------")
	log.Println(post.Body)
	for _, comment := range *comments {
		log.Println(comment.Body)
	}
}
