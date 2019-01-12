package main

import (
  //"fmt"
  "sync"
  "regexp"
  "strconv"
  "errors"
  "strings"
)

type Worker struct {
  id int
}

func (w *Worker) Start(waitGroup *sync.WaitGroup, queue chan string, result chan string){
  defer waitGroup.Done()

  requests := make(map[string]*Request)

  // Create regexp
  // TODO: If one of this regex err, quit program.
  // TODO: make these regex configurable
  urlRegex, _ := regexp.Compile(`\[(.+)\]`)
  timeRegex, _ := regexp.Compile(`(\d+)ms`)
  statusRegex, _ := regexp.Compile(`\|\s(\d+)\s.+`)

  for line := range queue {
    url, responseTime, responseCode, err := w.getRequestValues(line, urlRegex, timeRegex, statusRegex)
    if err == nil {
      // - check to see if Request with the same url exists in the requests map
      // - exist? update.
      // - not exist? create new Request and add to map.
      if req, ok := requests[url]; ok {
        req.Add(responseTime, responseCode)
      }else{
        result <- url // unique url
        requests[url] = NewRequest(url, responseTime, responseCode)
      }
    }
  }

  w.storeResults(requests)
}

/*
 * Extract url, response time, response code, isError from log line.
 * If one from those url, time and code is missing, it will return error=true.
 */
func (w *Worker) getRequestValues(
  line string,
  urlRegex *regexp.Regexp,
  timeRegex *regexp.Regexp,
  statusRegex *regexp.Regexp) (string, int, int, error) {

  if !urlRegex.MatchString(line) {
    return "", 0, 0, errors.New("url not found")
  }
  if !timeRegex.MatchString(line) {
    return "", 0, 0, errors.New("request time not found")
  }
  if !statusRegex.MatchString(line) {
    return "", 0, 0, errors.New("status code not found")
  }

  url := urlRegex.FindStringSubmatch(line)
  time := timeRegex.FindStringSubmatch(line)
  code := statusRegex.FindStringSubmatch(line)

  timeInt, _ := strconv.Atoi(time[1])
  codeInt, _ := strconv.Atoi(code[1])

  // extra query string cleanup
  var cleanedUrl string
  if strings.Contains(url[1], "?") {
    cleanedUrl = strings.Split(url[1], "?")[0]
  }else{
    cleanedUrl = url[1]
  }
  
  return cleanedUrl, timeInt, codeInt, nil
} 

/*
 * Save the results as CSV file with unique session ID as its base name.
 */
func (w *Worker) storeResults(results map[string]*Request) error {

  var rows [][]interface{}
  for _, v := range results {
    rows = append(rows, v.ToCsvData())
  }

  WriteCSV("filename.csv", RequestCsvHeader(), rows)
  return nil
}

func StartWorker(
  nWorker int,
  waitGroup *sync.WaitGroup,
  queue chan string,
  result chan string ){

  for id:=0; id < nWorker; id++ {
    w := &Worker{id: id}
    go w.Start(waitGroup, queue, result)
  }
}