package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func DownloadFile(url, destDir string) error {
	filename := filepath.Base(url)
	filePath := filepath.Join(destDir, filename) // .downloads/sample.txt
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	fmt.Println("Downloading", url)
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		_ = os.Remove(filePath)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		_ = os.Remove(filePath)
		return fmt.Errorf("bad status : %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Download %s took %s\n", filename, time.Since(start))
	return nil
}

func SequentialDownloader(urls []string, destDir string) error {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}
	start := time.Now()
	for _, url := range urls {
		if err := DownloadFile(url, destDir); err != nil {
			fmt.Println("Error downloading", url, err)
			continue
		}
	}
	fmt.Printf("Download %v took %s\n", urls, time.Since(start))
	return nil
}

type Result struct {
	URL      string
	FileName string
	Size     int64
	Duration time.Duration
	Error    error
}

func ConcurrentDownloader(urls []string, destDir string, maxConcurrent int) error {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}
	results := make(chan Result)
	var wg sync.WaitGroup
	limiter := make(chan struct{}, maxConcurrent)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			limiter <- struct{}{}
			defer func() {
				<-limiter
			}()

			start := time.Now()
			filename := filepath.Base(url)
			filePath := filepath.Join(destDir, filename)

			out, err := os.Create(filePath)
			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}
			defer out.Close()

			resp, err := http.Get(url)
			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}
			if resp.StatusCode != http.StatusOK {
				results <- Result{URL: url, Error: fmt.Errorf("bad status: %s", resp.Status)}
				return
			}
			size, err := io.Copy(out, resp.Body)
			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}
			timeSince := time.Since(start)
			results <- Result{URL: url, FileName: filename, Size: size, Duration: timeSince, Error: nil}
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var totalSize int64
	var errors []error
	start := time.Now()

	for result := range results {
		if result.Error != nil {
			errors = append(errors, result.Error)
		} else {
			totalSize += result.Size
			fmt.Printf("Downloaded %s (%d bytes) in %s\n", result.FileName, result.Size, result.Duration)
		}
	}
	startedSince := time.Since(start)
	fmt.Printf("Downloaded %d files, total size: %d bytes in %s\n", len(urls)-len(errors), totalSize, startedSince)
	if len(errors) > 0 {
		fmt.Printf("Errors: %v\n", errors)
	}
	return nil
}

func main() {
	// url := "https://picsum.photos/200"
	urls := []string{"https://placehold.co/600x400.png",
		"https://go.dev/images/go-logo-white.svg"}

	var err error // Declare the err variable
	// err = SequentialDownloader(urls, `.`) // Change to current directory
	err = ConcurrentDownloader(urls, `.`, 2) // Change to current directory
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("Done")
}
