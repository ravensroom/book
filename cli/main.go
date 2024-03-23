package main

import (
	"fmt"
	"github.com/ravensroom/book/cli/utils/env"
	"github.com/ravensroom/book/cli/utils/flag"
	"github.com/ravensroom/book/cli/utils/input"
	"github.com/ravensroom/book/pkg/agent"
	"github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

func main() {
	selectedModel := *flag.ModelFlag
	fmt.Printf("Using model <%s>...\n", selectedModel)
	agent := agent.NewAgent(
		env.OPENAI_API_KEY,
		agent.WithModel(agent.ModelType(selectedModel)),
	)
	var userQuestion *string
	for {
		fmt.Print("\n\033[1;33mUser:\033[0m ")
		input := input.GetUserInput()
		if input == "q!" {
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
		response := agent.ReadBotResponseStream(stream)
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
