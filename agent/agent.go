package agent

import (
	"github.com/ravensroom/code-hc/utils/env"
)

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
)

func GetBotResponseMessage(prompt string) string {
	client := openai.NewClient(env.OPENAI_API_KEY)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleUser, Content: prompt},
			},
		},
	)
	if err != nil {
		return "Error in openai chat completion " + err.Error()
	}
	return resp.Choices[0].Message.Content
}
