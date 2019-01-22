package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	// NumWorker is number of worker to run the import process
	NumWorker = 1
	// QueueSize is number of lines to be processed
	QueueSize = 500
)

// Import read the log file and save the result as a CSV file
func Import(filePath string, profile string) {
	if profile == "" {
		fmt.Println("[ERROR] Selected profile contains empty regex.")
		os.Exit(1)
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("[ERROR] %s", err)
		os.Exit(1)
	}

	fmt.Println("\nImport started with following configuration:")
	fmt.Printf("path=%v, profile=%v\n\n", filePath, profile)

	startTime := time.Now()
	importID := startTime.Unix() // a timestamp, will be used as csv filename.
	lineChan := make(chan string, QueueSize)
	resultChan := make(chan string)

	wg := &sync.WaitGroup{}
	wg.Add(NumWorker)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	StartWorker(importID, NumWorker, wg, lineChan, resultChan)

	reader := bufio.NewReader(file)
	go ReadLines(reader, lineChan)

	count := 0
	for range resultChan {
		fmt.Printf("\r> Importing %v unique endpoints...", count)
		count++
	}
	fmt.Printf("\nFinished in %.2fs, your import ID: %v\n\n", time.Since(startTime).Seconds(), importID)
	fmt.Printf("Now you can view the report with command: `rqmetric --view %v`\n\n", importID)
}
