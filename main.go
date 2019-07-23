package main

import (
	"log"
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

func getPlaceHolder() *PlaceHolder {
	config := tools.LoadConfig()
	return &PlaceHolder{url: config.Server.Url}
}

func printPosts(placeHolder *PlaceHolder, posts *[]Post) {
	commentsChan := make(chan *CommentsResult)
	defer close(commentsChan)
	go getComments(placeHolder, posts, commentsChan)

	for count := 0; count < len(*posts); count++ {
		commentResult := <-commentsChan
		if commentResult.err != nil {
			log.Fatal(commentResult.err)
		}
		log.Println(commentResult.comments)
		log.Println(count)
	}
}

func getComments(placeHolder *PlaceHolder, posts *[]Post, result chan<- (*CommentsResult)) {
	for _, post := range *posts {
		go placeHolder.GoGetPostComments(post.Id, result)
	}
}
