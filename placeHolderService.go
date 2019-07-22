package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	Url string
}

type Post struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

func GetPosts(server *Server) (*[]Post, error) {
	url := fmt.Sprintf("%s/posts", server.Url)
	resp, err := DoRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to get posts list from %s: %s", url, err)
	}
	body, err := ParseBodyToBytes(resp)
	if err != nil {
		return nil, err
	}
	var posts []Post
	if err := json.Unmarshal(*body, &posts); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal posts response")
	}
	return &posts, nil
}

type Comment struct {
	OostId int
	Id     int
	Name   string
	Email  string
	Body   string
}

func GetPostComments(server *Server, postId int) (*[]Comment, error) {
	url := fmt.Sprintf("%s/posts/%d/comments", server.Url, postId)
	resp, err := DoRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to get comments of posts %s from %s: %s", postId, url, err)
	}
	body, err := ParseBodyToBytes(resp)
	if err != nil {
		return nil, err
	}
	var comments []Comment
	if err := json.Unmarshal(*body, &comments); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal comment response")
	}
	return &comments, nil
}
