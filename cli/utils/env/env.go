package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/ravensroom/replica/cli/utils/input"
	"os"
	"strings"
)

var (
	OPENAI_API_KEY string
)

func init() {
	configFile := getConfigFilePath()
	if _, err := os.Stat(configFile); !os.IsNotExist(err) {
		content, _ := os.ReadFile(configFile)
		OPENAI_API_KEY = strings.TrimSpace(string(content))
	} else {
		fmt.Println("OPENAI_API_KEY not found in config file ~/.replicarc")
		err := godotenv.Load()
		if err != nil {
			fmt.Println(".env file not found. Reading system set environment variables")
		}
		OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
		if OPENAI_API_KEY == "" {
			fmt.Println("OPENAI_API_KEY environment variable not set")
		}
		AskAndSaveAPIKeyToConfig()
	}
}

func AskAndSaveAPIKeyToConfig() string {
	fmt.Print("Please (re)enter your API key: ")
	userInput := input.GetUserInput()
	OPENAI_API_KEY = strings.TrimSpace(userInput)
	configFile := getConfigFilePath()

	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		file, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			fmt.Printf("Failed to create config file: %s\n", err)
			os.Exit(1)
		}
		file.Close()
	} else if err != nil {
		fmt.Printf("Failed to check config file: %s\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(configFile, []byte(OPENAI_API_KEY), 0600)
	if err != nil {
		fmt.Printf("Failed to write to config file: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("API key saved to %s\n", configFile)
	return OPENAI_API_KEY
}

func getConfigFilePath() string {
	home, _ := os.UserHomeDir()
	return home + "/.replicarc"
}
