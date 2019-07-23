package tools

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func DoRequest(method string, url string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal("Failed to create request")
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}

func ParseBodyToBytes(resp *http.Response) (*[]byte, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response %s", err)
	}
	return &body, nil
}
