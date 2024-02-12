package main

import (
	"fmt"
	"github.com/ravensroom/code-hc/agent"
	"github.com/ravensroom/code-hc/agent/helper"
	"github.com/ravensroom/code-hc/utils/env"
	"github.com/ravensroom/code-hc/utils/flag"
	"github.com/ravensroom/code-hc/utils/input"
	"github.com/sashabaranov/go-openai"
	"os"
	"os/exec"
	"strings"
)

func main() {
	gitCommand := *flag.GitFlag
	gitOutput := getGitCommandOutput(gitCommand)
	fmt.Print("\033[1mEnter instruction for the bot: \033[0m")
	userInstruction := input.GetUserInput()
	agent := agent.NewAgent(userInstruction, gitOutput)
	var userQuestion *string
	processUserQuestionWithBot(agent, userQuestion)
	for {
		fmt.Print("\n\033[1;33mUser:\033[0m ")
		input := input.GetUserInput()
		if input == "q" {
			break
		}
		userQuestion = &input
		processUserQuestionWithBot(agent, userQuestion)
	}
}

func processUserQuestionWithBot(agent *agent.Agent, userQuestion *string) {
	if userQuestion != nil {
		agent.AddMessage(openai.ChatMessageRoleUser, *userQuestion)
	}
	stream, err := agent.GetBotResponseStream(userQuestion)
	if err != nil {
		fmt.Println("Error in getting bot response: ", err)
		if strings.Contains(err.Error(), "Invalid API key") {
			apiKey := env.AskAndSaveAPIKeyToConfig()
			agent.SetClient(openai.NewClient(apiKey))
			processUserQuestionWithBot(agent, userQuestion)
		} else {
			os.Exit(1)
		}
	}
	var botResponseMessage string
	fmt.Print("\n\033[1;32mBot:\033[0m ")
	for {
		response := helper.ReadBotResponseStream(stream)
		if response.Error != nil {
			fmt.Println("Error in reading bot response stream: ", err)
			os.Exit(1)
		}
		if response.EOF {
			agent.AddMessage(openai.ChatMessageRoleAssistant, botResponseMessage)
			fmt.Println()
			break
		}
		botResponseMessage += response.Chunk
		fmt.Print(response.Chunk)
	}
}

func getGitCommandOutput(gitCommand string) string {
	cmd := exec.Command("bash", "-c", gitCommand)
	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			fmt.Printf("%s command error: %s", gitCommand, exitErr.Stderr)
		}
		fmt.Printf("Error executing git command: %s", err)
		os.Exit(1)
	}
	if len(output) == 0 {
		fmt.Printf("Warning: No outpt found from command [%s]\n", gitCommand)
	}
	lineCount := 0
	for _, b := range output {
		if b == '\n' {
			lineCount++
		}
	}
	fmt.Printf("Output from [%s] command: %d lines\n", gitCommand, lineCount)
	return string(output)
}
