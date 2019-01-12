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

func main() {
  args := os.Args

  if len(args) < 2 {
    fmt.Println("Missing file name. Usage: `logmetric import ./production.log --profile=rails`")
    os.Exit(1)
  }

  file, err := os.Open(args[1])
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  startTime := time.Now()
  lineChan := make(chan string, QUEUE_SIZE)
  resultChan := make(chan string)

  wg := &sync.WaitGroup{}
  wg.Add(NUM_WORKER)

  go func (){
    wg.Wait()
    close(resultChan)
  }()

  StartWorker(NUM_WORKER, wg, lineChan, resultChan)

  reader := bufio.NewReader(file)
  go ReadLines(reader, lineChan)

  count := 0
  for range resultChan {
    fmt.Printf("\rProcessing %v unique endpoints...", count)
    count++
  }
  fmt.Printf("\nFinished in %v\n", time.Since(startTime))
}