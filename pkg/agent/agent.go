package agent

import (
	"context"
	"errors"
	openai "github.com/sashabaranov/go-openai"
	"io"
)

type Context struct {
	SystemInstruction string
	UserInstruction   string
	body              string
}

type Agent struct {
	client   *openai.Client
	model    ModelType
	contexts *[]Context
	messages *[]openai.ChatCompletionMessage
}

type AgentOptions func(*Agent)

func WithModel(model ModelType) AgentOptions {
	return func(a *Agent) {
		a.model = model
	}
}

func WithContexts(contexts *[]Context) AgentOptions {
	return func(a *Agent) {
		a.contexts = contexts
	}
}

func NewAgent(apiKey string, opts ...AgentOptions) *Agent {
	agent := &Agent{
		client: openai.NewClient(apiKey),
	}
	for _, opt := range opts {
		opt(agent)
	}
	if agent.model == "" {
		agent.model = GPT4Dot5
	}
	if agent.contexts != nil {
		agent.makeInitialMessages()
	}
	return agent
}

func (a *Agent) GetBotResponseStream(userQuestion *string) (*openai.ChatCompletionStream, error) {
	e := &openai.APIError{}
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model:    string(a.model),
		Messages: *a.messages,
		Stream:   true,
	}
	stream, err := a.client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		if errors.As(err, &e) {
			switch e.HTTPStatusCode {
			case 401:
				return nil, errors.New("Invalid API key")
			case 429:
				return nil, errors.New("Rate limit exceeded")
			case 500:
				return nil, errors.New("Openai server error")
			}
		}
		return nil, err
	}
	return stream, nil
}

type BotResponse struct {
	Chunk string
	EOF   bool
	Error error
}

func (a *Agent) ReadBotResponseStream(stream *openai.ChatCompletionStream) BotResponse {
	response, err := stream.Recv()
	if errors.Is(err, io.EOF) {
		return BotResponse{
			Chunk: "",
			EOF:   true,
			Error: nil,
		}
	}
	if err != nil {
		return BotResponse{
			Chunk: "",
			EOF:   false,
			Error: err,
		}
	}
	return BotResponse{
		Chunk: response.Choices[0].Delta.Content,
		EOF:   false,
		Error: nil,
	}
}

func (a *Agent) makeInitialMessages() {
	if a.contexts == nil {
		return
	}
	for _, context := range *a.contexts {
		a.AddMessage(openai.ChatMessageRoleAssistant, context.SystemInstruction)
		a.AddMessage(openai.ChatMessageRoleUser, context.UserInstruction)
		a.AddMessage(openai.ChatMessageRoleUser, context.body)
	}
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

func (a *Agent) SetClient(client *openai.Client) {
	a.client = client
}
