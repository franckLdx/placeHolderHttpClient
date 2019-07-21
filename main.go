package main

import (
	"fmt"
	"log"
)

func main() {
	data, err := getPosts()
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range *data {
		fmt.Println("-----------------------")
		id := fmt.Sprintf("%.0f", post["id"])
		c, err := getComments(id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(c)
	}
}

func getPosts() (*[]map[string]interface{}, error) {
	resp, err := DoRequest("GET", "https://jsonplaceholder.typicode.com/posts", nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to get posts list %s", err)
	}
	return ParseBodyToUnknownArray(resp)
}

func getComments(id string) (*[]map[string]interface{}, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s/comments", id)
	resp, err := DoRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to get comments of posts %s: %s", id, err)
	}
	return ParseBodyToUnknownArray(resp)
}
