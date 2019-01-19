package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

// WriteCSV write 2 dimensional array into a csv file.
func WriteCSV(fileName string, header []string, data [][]string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Panic("Unable to write results to the disk!")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write the header
	writer.Write(header)

	// write rows
	for _, row := range data {
		writer.Write(row)
	}
}

// ReadCSV read csv file and return a 2 dimensional array.
func ReadCSV(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Panic("Unable to read file from the disk!")
	}
	defer file.Close()

	var rows [][]string
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		rows = append(rows, row)
	}
	return rows
}
