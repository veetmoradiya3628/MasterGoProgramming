package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

type PostsClient struct {
	httpClient HTTPClient
	baseURL    string
}

func NewPostsClient(httpClient HTTPClient, baseURL string) *PostsClient {
	return &PostsClient{
		httpClient: httpClient,
		baseURL:    baseURL,
	}
}

type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

type PostResponse struct {
	Posts []Post `json:"posts"`
	Total int    `json:"total"`
}

func (c *PostsClient) FetchPosts(limit int) ([]Post, error) {
	url := fmt.Sprintf("%s/posts?limit=%d", c.baseURL, limit)
	resp, err := c.httpClient.Get(url) // depend on interface
	if err != nil {
		return nil, fmt.Errorf("failed to fetch posts : %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	var postResp PostResponse
	if err := json.NewDecoder(resp.Body).Decode(&postResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return postResp.Posts, nil
}

func main() {

	baseUrl := "https://dummyjson.com"
	postClient := NewPostsClient(http.DefaultClient, baseUrl)
	posts, err := postClient.FetchPosts(3)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	for _, post := range posts {
		fmt.Printf("[%d] %s\n", post.ID, post.Title)
	}
}
