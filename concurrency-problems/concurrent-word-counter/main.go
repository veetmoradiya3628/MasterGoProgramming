package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

// countWords processes a single file (shared by both methods)
func countWords(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return wordCount, nil
}

// sequentialRead loops through files 1 to 20 one by one
func sequentialRead(basePath string) int {
	grandTotal := 0

	for i := 1; i <= 20; i++ {
		filePath := fmt.Sprintf("%s/file_%d.txt", basePath, i)

		fileStart := time.Now()
		count, err := countWords(filePath)
		fileDuration := time.Since(fileStart)

		if err != nil {
			fmt.Printf("[Seq] Warning: Failed to read %s: %v (Took: %v)\n", filePath, err, fileDuration)
			continue
		}

		fmt.Printf("[Seq] File file_%d.txt | Words: %d | Time: %v\n", i, count, fileDuration)
		grandTotal += count
	}

	return grandTotal
}

// parallelRead loops through files 1 to 20 concurrently
func parallelRead(basePath string) int {
	var wg sync.WaitGroup
	resultsChan := make(chan int, 20)

	for i := 1; i <= 20; i++ {
		wg.Add(1)

		go func(fileNum int) {
			defer wg.Done()

			filePath := fmt.Sprintf("%s/file_%d.txt", basePath, fileNum)

			fileStart := time.Now()
			count, err := countWords(filePath)
			fileDuration := time.Since(fileStart)

			if err != nil {
				fmt.Printf("[Par] Warning: Failed to read %s: %v (Took: %v)\n", filePath, err, fileDuration)
				resultsChan <- 0
				return
			}

			fmt.Printf("[Par] File file_%d.txt | Words: %d | Time: %v\n", fileNum, count, fileDuration)
			resultsChan <- count
		}(i)
	}

	wg.Wait()
	close(resultsChan)

	grandTotal := 0
	for count := range resultsChan {
		grandTotal += count
	}

	return grandTotal
}

func main() {
	baseFolder := "inputs"

	fmt.Println("--- Starting Sequential Execution ---")
	seqStart := time.Now()
	seqTotal := sequentialRead(baseFolder)
	seqDuration := time.Since(seqStart)

	fmt.Println("\n--- Starting Parallel Execution ---")
	parStart := time.Now()
	parTotal := parallelRead(baseFolder)
	parDuration := time.Since(parStart)

	fmt.Println("\n==============================================")
	fmt.Println("                FINAL RESULTS                 ")
	fmt.Println("==============================================")
	fmt.Printf("Sequential Total : %d words | Time: %v\n", seqTotal, seqDuration)
	fmt.Printf("Parallel Total   : %d words | Time: %v\n", parTotal, parDuration)
	fmt.Println("==============================================")
}
