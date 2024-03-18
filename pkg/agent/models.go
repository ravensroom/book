package agent

import (
	openai "github.com/sashabaranov/go-openai"
)

type ModelType string

const (
	GPT3Dot5 ModelType = openai.GPT3Dot5Turbo0125
	GPT4Dot5 ModelType = openai.GPT4Turbo0125
)
