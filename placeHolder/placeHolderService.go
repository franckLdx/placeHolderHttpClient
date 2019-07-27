package placeHolder

import (
	"encoding/json"
	"fmt"
	"placeHolderHttpClient/tools"
)

type Server struct {
	Url string
}

func (server *Server) GetPosts() (*Posts, error) {
	url := fmt.Sprintf("%s/posts", server.Url)
	body, err := getBytes("GET", url)
	if err != nil {
		return nil, fmt.Errorf("Failed to get posts list from %s: %s", url, err)
	}
	var posts Posts
	if err := json.Unmarshal(*body, &posts); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal posts response")
	}
	return &posts, nil
}

func (server *Server) GetComments(postId int) (*Comments, error) {
	url := fmt.Sprintf("%s/posts/%d/comments", server.Url, postId)
	body, err := getBytes("GET", url)
	if err != nil {
		return nil, fmt.Errorf("Failed to get comments of posts %d from %s: %s", postId, url, err)
	}
	var comments Comments
	if err := json.Unmarshal(*body, &comments); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal comment response")
	}
	return &comments, nil
}

func (server *Server) GoGetComments(postId int, result chan<- *CommentsResult) {
	comments, err := server.GetComments(postId)
	result <- &CommentsResult{postId, *comments, err}
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
