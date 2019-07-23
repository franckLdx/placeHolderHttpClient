package main

import (
	"encoding/json"
	"fmt"
	"placeHolderHttpClient/tools"
	"placeHolderHttpClient/types"
)

func GetPosts(server *types.Server) (*[]types.Post, error) {
	url := fmt.Sprintf("%s/posts", server.Url)
	resp, err := tools.DoRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to get posts list from %s: %s", url, err)
	}
	body, err := tools.ParseBodyToBytes(resp)
	if err != nil {
		return nil, err
	}
	var posts []types.Post
	if err := json.Unmarshal(*body, &posts); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal posts response")
	}
	return &posts, nil
}

func GetPostComments(server *types.Server, postId int) (*[]types.Comment, error) {
	url := fmt.Sprintf("%s/posts/%d/comments", server.Url, postId)
	resp, err := tools.DoRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to get comments of posts %s from %s: %s", postId, url, err)
	}
	body, err := tools.ParseBodyToBytes(resp)
	if err != nil {
		return nil, err
	}
	var comments []types.Comment
	if err := json.Unmarshal(*body, &comments); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal comment response")
	}
	return &comments, nil
}
