package main

import (
  "fmt"
  "flag"
)

func printUsage() {
  fmt.Printf("\n== [%v %v] ==\n", AppName, AppVersion)
  fmt.Println("\nUsage examples:")
  fmt.Println("Import log file  =>\trqmetric -import=production.log -profile=rails")
  fmt.Println("Serve the report =>\trqmetric -serve=123456")
  fmt.Println("Params help      =>\trqmetric -h\n")
}

func main() {
  filePath := flag.String("import", "", "Path to the file that will be imported")
  profile := flag.String("profile", "rails", "Log profile to be use to parse the log line")
  serveId := flag.String("serve", "", "Import ID to be served")

  flag.Parse()

  if *filePath != "" {
    Import(*filePath, *profile)
  } else if *serveId != "" {
    Serve(*serveId)
  } else {
    printUsage()
  }
}