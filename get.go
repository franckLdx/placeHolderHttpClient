package main

import (
	"log"
	"placeHolderHttpClient/placeHolder"
)

func GetPosts(server *placeHolder.Server) *placeHolder.Posts {
	posts, err := server.GetPosts()
	if err != nil {
		log.Fatal(err)
	}
	return posts
}

func GetComments(server *placeHolder.Server, posts *placeHolder.Posts) *placeHolder.CommentsMap {
	commentsChan := make(chan *placeHolder.CommentsResult)
	defer close(commentsChan)
	go startGetComments(server, posts, commentsChan)
	return fillCommentsMap(len(*posts), commentsChan)
}

func startGetComments(server *placeHolder.Server, posts *placeHolder.Posts, commentsChan chan<- *placeHolder.CommentsResult) {
	for _, post := range *posts {
		go server.GoGetComments(post.Id, commentsChan)
	}
}

func fillCommentsMap(postsCount int, commentsChan <-chan *placeHolder.CommentsResult) *placeHolder.CommentsMap {
	comments := make(placeHolder.CommentsMap)
	for count := 0; count < postsCount; count++ {
		commentsResult := <-commentsChan
		if commentsResult.Err != nil {
			log.Fatal("Failed to get a comment ", commentsResult.Err)
		}
		if comments[commentsResult.PostId] != nil {
			log.Fatal("Halready have a comments list for post ", commentsResult.PostId)
		}
		comments[commentsResult.PostId] = commentsResult.Comments
	}
	return &comments
}
