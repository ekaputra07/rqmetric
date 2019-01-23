package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/viper"
)

func printUsage() {
	fmt.Printf("\n== [%v %v - %v] ==\n", AppName, AppVersion, AppRepo)
	fmt.Println("\nUsage examples:")
	fmt.Println("Create profile config =>\trqmetric -init")
	fmt.Println("Import log file       =>\trqmetric -import production.log -profile rails")
	fmt.Println("View the report       =>\trqmetric -view 123456")
	fmt.Println("Params help           =>\trqmetric -h")
}

func main() {

	viper.SetConfigType("yaml")
	viper.SetConfigName(".rqmetric")
	viper.AddConfigPath("$HOME")

	initProfile := flag.Bool("init", false, fmt.Sprintf("Create initial profile config in %s/.rqmetric.yml", os.Getenv("HOME")))
	filePath := flag.String("import", "", "Path to the file that will be imported")
	profile := flag.String("profile", "default", "Log profile to be use to parse the log line")
	viewImportID := flag.String("view", "", "Import ID to be viewed")

	flag.Parse()

	if *initProfile {
		InitConfig()
		return
	}

	if *filePath != "" {
		LoadConfig()
		regex := viper.GetString(*profile)

		if regex == "" {
			fmt.Println("[ERROR] Selected profile does not exist or empty.")
			os.Exit(1)
		}

		re, err := regexp.Compile(regex)
		if err != nil {
			fmt.Printf("[ERROR] %s in:\n%s\n", err, regex)
			os.Exit(1)
		}
		Import(*filePath, *profile, re)

	} else if *viewImportID != "" {
		StartViewer(*viewImportID)
	} else {
		printUsage()
	}
}
