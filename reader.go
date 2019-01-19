package main

import (
	"bufio"
	"strings"
)

// readLine read and return a single line
func readLine(r *bufio.Reader) (string, error) {
	var (
		isPrefix = true
		err      error
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

// ReadLines read all lines of the log files and pass the results to line channel
func ReadLines(reader *bufio.Reader, lineChan chan string) {
	s, e := readLine(reader)

	for e == nil {
		if strings.Contains(s, "Completed") {
			lineChan <- s
		}
		s, e = readLine(reader)
	}

	close(lineChan)
}
