package main

import (
	"flag"
	"fmt"
	"github.com/ravensroom/code-hc/agent"
)

// Define cmd line flags
var (
	messageFlag = flag.String("m", "", "General context of the commit")
)

func main() {
	flag.Parse()
	userMessage := *messageFlag
	if userMessage == "" {
		fmt.Println("Error: -m flag is required")
		return
	}
	fmt.Println(agent.GetBotResponseMessage(userMessage))
}
