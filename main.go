package main

import (
	"bufio"
	"fmt"
	"github.com/ravensroom/code-hc/agent"
	"github.com/ravensroom/code-hc/agent/helper"
	"github.com/sashabaranov/go-openai"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Print("Background information for the bot: ")
	userInstruction := getUserInput()

	fmt.Print("Enter the git diff command: ")
	gitDiffCmd := getUserInput()
	if gitDiffCmd == "" {
		gitDiffCmd = "git diff"
	}
	diffResult, err := getGitDiff(gitDiffCmd)
	for err != nil {
		fmt.Println("Error in getting git diff: ", err)
		fmt.Print("Please re-enter the command or enter q to quit: ")
		gitDiffCmd = getUserInput()
		if gitDiffCmd == "q" {
			os.Exit(0)
		} else {
			diffResult, err = getGitDiff(gitDiffCmd)
		}
	}
	agent := agent.NewAgent(userInstruction, diffResult)
	fmt.Println("\nYour instruction and git diff have been uploaded for bot to process. Fetching bot reply...")
	var userQuestion *string
	processUserQuestionWithBot(agent, userQuestion)
	for {
		fmt.Print("\nUser: ")
		input := getUserInput()
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
		os.Exit(1)
	}
	var botResponseMessage string
	fmt.Print("\nBot: ")
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

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func getGitDiff(diffCmd string) (string, error) {
	if !strings.HasPrefix(diffCmd, "git diff") {
		return "", fmt.Errorf("Invalid git diff command")
	}
	cmd := exec.Command("bash", "-c", diffCmd)
	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return "", fmt.Errorf("Git diff command error: %s", exitErr.Stderr)
		}
		return "", fmt.Errorf("Error executing git diff command: %s", err)
	}
	return string(output), nil
}
