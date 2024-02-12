package agent

import (
	"context"
	"github.com/ravensroom/code-hc/utils/env"
	openai "github.com/sashabaranov/go-openai"
)

type Agent struct {
	client      *openai.Client
	model       string
	initContext string
	messages    *[]openai.ChatCompletionMessage
}

func NewAgent(userInstruction string, gitDiff string) *Agent {
	instruction := "The user is asking you do perform a code health check. He's likely reviewing a pull request or about to commit his code and he wants to make sure that the code is healthy and generally following best practices. He has provided the following background information: {" + userInstruction + "} and the git diff content: " + gitDiff + ". Please provide your suggestions to the user."

	return &Agent{
		client:      openai.NewClient(env.OPENAI_API_KEY),
		model:       openai.GPT4,
		initContext: instruction,
		messages: &[]openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: instruction,
			},
		},
	}
}

func (a *Agent) GetBotResponseMessage(userQuestion *string) string {
	if userQuestion != nil {
		*a.messages = append(*a.messages, openai.ChatCompletionMessage{
			Role:    "user",
			Content: *userQuestion,
		})
	}

	resp, err := a.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    a.model,
			Messages: *a.messages,
		},
	)
	if err != nil {
		return "Error in openai chat completion " + err.Error()
	}
	responseMessage := resp.Choices[0].Message
	*a.messages = append(*a.messages, responseMessage)
	return responseMessage.Content
}
