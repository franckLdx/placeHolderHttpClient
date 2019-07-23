package main

import (
	"log"
	"placeHolderHttpClient/tools"
)

func main() {
	server := getServerConfig()
	posts, err := GetPosts(server)
	if err != nil {
		log.Fatal(err)
	}
	printPosts(posts)
}

func getServerConfig() *Server {
	config := tools.LoadConfig()
	return &config.Server
}

func printPosts() {
	for _, post := range *data {
		log.Println("-----------------------")
		log.Println("id: ", post.Id)
		log.Println("title: ", post.Title)
		c, err := GetPostComments(server, post.Id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(c)
	}
}
