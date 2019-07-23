package main

import (
	"encoding/json"
	"fmt"
	"placeHolderHttpClient/tools"
)

type PlaceHolder struct {
	url string
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

func (placeHolder *PlaceHolder) GetPosts() (*[]Post, error) {
	url := fmt.Sprintf("%s/posts", placeHolder.url)
	resp, err := tools.DoRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to get posts list from %s: %s", url, err)
	}
	body, err := tools.ParseBodyToBytes(resp)
	if err != nil {
		return nil, err
	}
	var posts []Post
	if err := json.Unmarshal(*body, &posts); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal posts response")
	}
	return &posts, nil
}

func (placeHolder *PlaceHolder) GetPostComments(postId int) (*[]Comment, error) {
	url := fmt.Sprintf("%s/posts/%d/comments", placeHolder.url, postId)
	resp, err := tools.DoRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to get comments of posts %s from %s: %s", postId, url, err)
	}
	body, err := tools.ParseBodyToBytes(resp)
	if err != nil {
		return nil, err
	}
	var comments []Comment
	if err := json.Unmarshal(*body, &comments); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal comment response")
	}
	return &comments, nil
}

type CommentsResult struct {
	comments []Comment
	err      error
}

func (placeHolder *PlaceHolder) GoGetPostComments(postId int, result chan<- (*CommentsResult)) {
	comments, err := placeHolder.GetPostComments(postId)
	result <- &CommentsResult{*comments, err}
}
