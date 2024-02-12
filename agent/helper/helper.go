package helper

import (
	"errors"
	openai "github.com/sashabaranov/go-openai"
	"io"
)

type BotResponse struct {
	Chunk string
	EOF   bool
	Error error
}

func ReadBotResponseStream(stream *openai.ChatCompletionStream) BotResponse {
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
