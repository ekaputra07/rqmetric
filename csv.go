package main

import (
  "os"
  "encoding/csv"
  "log"
)

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