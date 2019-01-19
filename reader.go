package main

import (
	"bufio"
	"strings"
)

func readLine(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

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
