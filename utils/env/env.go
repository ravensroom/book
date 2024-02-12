package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/ravensroom/code-hc/utils/input"
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
		fmt.Println("OPENAI_API_KEY not found in config file ~/.codehcrc")
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
	info, _ := os.Stat(configFile)
	permission := info.Mode().Perm()
	if permission|0200 != 0 {
		fmt.Println("No write permission to config file ~/.codehcrc")
		os.Exit(1)
	}
	_ = os.WriteFile(configFile, []byte(OPENAI_API_KEY), 0600)
	return OPENAI_API_KEY
}

func getConfigFilePath() string {
	home, _ := os.UserHomeDir()
	return home + "/.codehcrc"
}
