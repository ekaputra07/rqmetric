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
	fmt.Println("Import log file       =>\trqmetric -import production.log -profile default -max-url-parts-length=20")
	fmt.Println("View the report       =>\trqmetric -view rqmetric_output_123456.csv")
	fmt.Println("Params help           =>\trqmetric -h")
}

func main() {

	viper.SetConfigType("yaml")
	viper.SetConfigName(".rqmetric")
	viper.AddConfigPath("$HOME")

	initProfile := flag.Bool("init", false, fmt.Sprintf("Create initial profile config in %s/.rqmetric.yml", os.Getenv("HOME")))
	filePath := flag.String("import", "", "Path to the file that will be imported")
	profile := flag.String("profile", "default", "Log profile to be use to parse the log line")
	maxURLPartsLength := flag.Int64("max-url-parts-length", 0, "Maximum length of url parts (separated by '/') that shouldn't treated as a parameter")
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
		Import(*filePath, *profile, re, *maxURLPartsLength)

	} else if *viewImportID != "" {
		StartViewer(*viewImportID)
	} else {
		printUsage()
	}
}
