package main

import (
	"flag"
	"fmt"
)

func printUsage() {
	fmt.Printf("\n== [%v %v - %v] ==\n", AppName, AppVersion, AppRepo)
	fmt.Println("\nUsage examples:")
	fmt.Println("Import log file  =>\trqmetric --import production.log --profile rails")
	fmt.Println("View the report  =>\trqmetric --view 123456")
	fmt.Println("Serve the report =>\trqmetric --serve 123456 --port 8080")
	fmt.Println("Params help      =>\trqmetric -h")
}

func main() {
	filePath := flag.String("import", "", "Path to the file that will be imported")
	profile := flag.String("profile", "rails", "Log profile to be use to parse the log line")
	serveImportID := flag.String("serve", "", "Import ID to be served")
	servePort := flag.String("port", "5000", "Port number to bind the http service, must be used with --serve param")
	viewImportID := flag.String("view", "", "Import ID to be viewed")

	flag.Parse()

	if *filePath != "" {
		Import(*filePath, *profile)
	} else if *serveImportID != "" {
		Serve(*serveImportID, *servePort)
	} else if *viewImportID != "" {
		StartViewer(*viewImportID)
	} else {
		printUsage()
	}
}
