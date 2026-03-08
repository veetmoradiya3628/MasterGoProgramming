package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

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

func fetchPosts(limit int) ([]Post, error) {
	url := fmt.Sprintf("https://dummyjson.com/posts?limit=%d", limit)
	resp, err := http.Get(url) // hard dependency
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
	posts, err := fetchPosts(2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	for _, post := range posts {
		fmt.Printf("[%d] %s\n", post.ID, post.Title)
	}
}
