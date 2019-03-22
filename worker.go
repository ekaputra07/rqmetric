package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

// Worker is an object that keep tracks of the processing log line.
type Worker struct {
	importID  int64
	re        *regexp.Regexp
	maxURLPartsLength int64
	waitGroup *sync.WaitGroup
	queue     chan string
}

// Start will listen for result channel and process every text passed to it.
func (w *Worker) Start(result chan string) {
	defer w.waitGroup.Done()

	requests := make(map[string]*Request)

	for line := range w.queue {
		url, responseTime, responseCode := w.getRequestValues(line)
		// - check to see if Request with the same url exists in the requests map
		// - exist? update.
		// - not exist? create new Request and add to map.
		if req, ok := requests[url]; ok {
			req.Add(responseTime, responseCode)
		} else {
			result <- url // unique url
			requests[url] = NewRequest(url, responseTime, responseCode)
		}
	}

	w.storeResults(requests)
}

// getRequestValues extracts url, response time, response code, isError from log line.
// If one from those url, time and code is missing, it will return error=true.
func (w *Worker) getRequestValues(line string) (string, int, int) {

	match := w.re.FindStringSubmatch(line)
	names := w.re.SubexpNames()
	values := make(map[string]string)

	for i, value := range match {
		values[names[i]] = value
	}

	url := values["url"]
	respTime, _ := strconv.Atoi(values["time"])
	respCode, _ := strconv.Atoi(values["code"])

	// extra query string cleanup
	var cleanedURL string
	if strings.Contains(url, "?") {
		cleanedURL = strings.Split(url, "?")[0]
	} else {
		cleanedURL = url
	}

	// If the maxURLPartsLength == 0, return the url as is.
	if w.maxURLPartsLength == 0 {
		return cleanedURL, respTime, respCode
	}

	// remove params from url
	// - split the path
	// - check each parts length
	// - if its > maxURLPartsLength, replace it with <param>
	parts := strings.Split(cleanedURL, "/")
	var newParts []string
	for _, p := range parts {
		if len(p) > int(w.maxURLPartsLength) {
			newParts = append(newParts, fmt.Sprintf("<param:%v>", len(p)))
		} else {
			newParts = append(newParts, p)
		}
	}
	return strings.Join(newParts[:], "/"), respTime, respCode
}

// storeResults saves the results as CSV file with unique session ID as its base name.
func (w *Worker) storeResults(results map[string]*Request) {

	var rows [][]string
	for _, v := range results {
		if v.Count < 10 {
			continue
		}
		rows = append(rows, v.ToCsvData())
	}

	WriteCSV(fmt.Sprintf("rqmetric_output_%v.csv", w.importID), RequestCsvHeader(), rows)
}

// StartWorker starts the worker goroutines.
func StartWorker(
	importID int64,
	re *regexp.Regexp,
	maxURLPartsLength int64,
	waitGroup *sync.WaitGroup,
	queue chan string,
	result chan string) {

	w := &Worker{importID, re, maxURLPartsLength, waitGroup, queue}
	go w.Start(result)
}
