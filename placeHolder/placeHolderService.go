package placeHolder

import (
	"encoding/json"
	"fmt"
	"placeHolderHttpClient/tools"
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

type Comment struct {
	PostId int
	Id     int
	Name   string
	Email  string
	Body   string
}

func (server *Server) GetPosts() (*[]Post, error) {
	url := fmt.Sprintf("%s/posts", server.Url)
	body, err := getBytes("GET", url)
	if err != nil {
		return nil, fmt.Errorf("Failed to get posts list from %s: %s", url, err)
	}
	var posts []Post
	if err := json.Unmarshal(*body, &posts); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal posts response")
	}
	return &posts, nil
}

func (server *Server) GetComments(postId int) (*[]Comment, error) {
	url := fmt.Sprintf("%s/posts/%d/comments", server.Url, postId)
	body, err := getBytes("GET", url)
	if err != nil {
		return nil, fmt.Errorf("Failed to get comments of posts %s from %s: %s", postId, url, err)
	}
	var comments []Comment
	if err := json.Unmarshal(*body, &comments); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal comment response")
	}
	return &comments, nil
}

type CommentsResult struct {
	Comments []Comment
	Err      error
}

func (server *Server) GoGetComments(postId int, result chan<- (*CommentsResult)) {
	comments, err := server.GetComments(postId)
	result <- &CommentsResult{*comments, err}
}

func getBytes(method, url string) (*[]byte, error) {
	resp, err := tools.DoRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	body, err := tools.ParseBodyToBytes(resp)
	if err != nil {
		return nil, err
	}
	return body, nil
}
