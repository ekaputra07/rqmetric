package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

// LoadConfig try to load config file and exit if config file not found.
func LoadConfig() {
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("[ERROR] %s\n\n", err)
		fmt.Println("If you haven't create it yet, create one with:")
		fmt.Println("rqmetric -init")
		os.Exit(1)
	}
}

// InitConfig create initial config file on specified path.
// The `default` regex will match the following log line:
// `Completed in 19ms (View: 0, DB: 5) | 200 OK [https://example.com/about.html]`
func InitConfig() {
	exampleProfile := []byte(`default: 'Completed in (?P<time>\d+)ms .* (?P<code>\d+) .* \[(?P<url>.+)\]'`)

	err := ioutil.WriteFile(fmt.Sprintf("%s/.rqmetric.yml", os.Getenv("HOME")), exampleProfile, 0644)
	if err != nil {
		fmt.Println("[ERROR] Unable to create config file, please make sure you have a sufficient permission.")
		os.Exit(1)
	}

	fmt.Printf("[OK] Config file created: %s/.rqmetric.yml\n", os.Getenv("HOME"))
	fmt.Println("Now you can edit the default profile or add a new one.")
	os.Exit(0)
}
