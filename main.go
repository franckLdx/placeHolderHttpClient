package main

import (
	"log"
	"placeHolderHttpClient/placeHolder"
	"placeHolderHttpClient/tools"
)

func main() {
	server := getplaceHolder()

	posts := GetPosts(server)

	comments := GetComments(server, posts)

	printPostsWithComments(posts, comments)

	log.Println("Bye Bye")
}

func getplaceHolder() *placeHolder.Server {
	config := tools.LoadConfig()
	return &placeHolder.Server{Url: config.Server.Url}
}

func printPostsWithComments(posts *placeHolder.Posts, comments *placeHolder.CommentsMap) {
	for _, post := range *posts {
		postComments := (*comments)[post.Id]
		if postComments == nil {
			log.Fatal("Post %d has no comments", post.Id)
		}
		printPostWithComments(&post, &postComments)
	}

}

func printPostWithComments(post *placeHolder.Post, comments *placeHolder.Comments) {
	log.Println("----------------------------------------------------")
	log.Println(post.Body)
	for _, comment := range *comments {
		log.Println(comment.Body)
	}
}
