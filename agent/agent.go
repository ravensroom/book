package agent

import (
	"context"
	"errors"
	"github.com/ravensroom/code-hc/utils/env"
	openai "github.com/sashabaranov/go-openai"
)

type Agent struct {
	client      *openai.Client
	model       string
	initContext string
	messages    *[]openai.ChatCompletionMessage
}

func NewAgent(userInstruction string, gitOutput string) *Agent {
	instruction := "The user is asking you do perform a code health check. He's likely reviewing a pull request or about to commit his code and he wants to make sure that the code is healthy and generally following best practices. He has provided the following instruction: {" + userInstruction + "} and the git show/diff result: {" + gitOutput + "}. Please provide your suggestions to the user."

	agent := &Agent{
		client:      openai.NewClient(env.OPENAI_API_KEY),
		model:       openai.GPT4,
		initContext: instruction,
	}
	agent.AddMessage(openai.ChatMessageRoleSystem, instruction)
	return agent
}

func (a *Agent) GetBotResponseStream(userQuestion *string) (*openai.ChatCompletionStream, error) {
	e := &openai.APIError{}
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model:    a.model,
		Messages: *a.messages,
		Stream:   true,
	}
	stream, err := a.client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		if errors.As(err, &e) {
			switch e.HTTPStatusCode {
			case 401:
				return nil, errors.New("Invalid API key (do not retry)")
			case 429:
				return nil, errors.New("Rate limit exceeded (wait and retry)")
			case 500:
				return nil, errors.New("openai server error (wait and retry)")
			}
		}
		return nil, err
	}
	return stream, nil
}

func (a *Agent) AddMessage(role string, content string) {
	if a.messages == nil {
		a.messages = &[]openai.ChatCompletionMessage{}
	}
	*a.messages = append(*a.messages, openai.ChatCompletionMessage{
		Role:    role,
		Content: content,
	})
}
