package main

import (
  "fmt"
  "os"
  "bufio"
  "time"
  "sync"
)

const (
  NUM_WORKER = 1
  QUEUE_SIZE = 500
)

func Import(filePath string, profile string) {
  file, err := os.Open(filePath)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Println("\nImport started with following configuration:")
  fmt.Printf("path=%v, profile=%v\n\n", filePath, profile)

  startTime := time.Now()
  sessionId := startTime.Unix() // a timestamp, will be used as csv filename.
  lineChan := make(chan string, QUEUE_SIZE)
  resultChan := make(chan string)

  wg := &sync.WaitGroup{}
  wg.Add(NUM_WORKER)

  go func (){
    wg.Wait()
    close(resultChan)
  }()

  StartWorker(sessionId, NUM_WORKER, wg, lineChan, resultChan)

  reader := bufio.NewReader(file)
  go ReadLines(reader, lineChan)

  count := 0
  for range resultChan {
    fmt.Printf("\r> Importing %v unique endpoints...", count)
    count++
  }
  fmt.Printf("\nFinished in %.2fs, your import ID: %v\n\n", time.Since(startTime).Seconds(), sessionId)
  fmt.Printf("Now you can view the report with command: `rqmetric -serve=%v`\n\n", sessionId)
}