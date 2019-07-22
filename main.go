package main

import (
	"log"
)

func main() {
	server := getServerConfig()
	data, err := GetPosts(server)
	if err != nil {
		log.Fatal(err)
	}
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

func getServerConfig() *Server {
	config := LoadConfig()
	return &config.Server
}
