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
		fmt.Printf("\nIt will create a config file: %s/.rqmetric.json", os.Getenv("HOME"))
		os.Exit(1)
	}
}

// InitConfig create initial config file on specified path.
func InitConfig() {
	exampleProfile := []byte(`{
    "default": "rEGEX"
  }`)

	err := ioutil.WriteFile(fmt.Sprintf("%s/.rqmetric.json", os.Getenv("HOME")), exampleProfile, 0644)
	if err != nil {
		fmt.Println("[ERROR] Unable to create config file, please make sure you have a sufficient permission.")
		os.Exit(1)
	}

	fmt.Printf("[OK] Profile config created: %s/.rqmetric.json\n", os.Getenv("HOME"))
	fmt.Println("Now you can edit the default profile or add a new one.")
	os.Exit(0)
}
