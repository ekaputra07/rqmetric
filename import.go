package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"
	"time"
)

const queueSize = 500

// Import read the log file and save the result as a CSV file
func Import(filePath, profile string, re *regexp.Regexp) {

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		os.Exit(1)
	}

	fmt.Println("\nImport started with following configuration:")
	fmt.Printf("path=%v, profile=%v\n\n", filePath, profile)

	startTime := time.Now()
	importID := startTime.Unix() // a timestamp, will be used as csv filename.
	lineChan := make(chan string, queueSize)
	resultChan := make(chan string)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	StartWorker(importID, re, wg, lineChan, resultChan)

	reader := bufio.NewReader(file)
	go ReadLines(reader, re, lineChan)

	count := 1
	for range resultChan {
		fmt.Printf("\r> Importing %v unique endpoints...", count)
		count++
	}
	fmt.Printf("\nFinished in %.2fs, your import ID: %v\n\n", time.Since(startTime).Seconds(), importID)
	fmt.Printf("Now you can view the report with command: `rqmetric -view %v`\n\n", importID)
}
