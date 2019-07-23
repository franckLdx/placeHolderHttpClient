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
}

func getPlaceHolder() *PlaceHolder {
	config := tools.LoadConfig()
	return &PlaceHolder{url: config.Server.Url}
}

func printPosts(placeHolder *PlaceHolder, posts *[]Post) {
	for _, post := range *posts {
		log.Println("-----------------------")
		log.Println("id: ", post.Id)
		log.Println("title: ", post.Title)
		c, err := placeHolder.GetPostComments(post.Id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(c)
	}
}
