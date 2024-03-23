package flag

import (
	"flag"
	"fmt"
	"github.com/ravensroom/book/pkg/agent"
)

var (
	ModelFlag = flag.String("m", "1",
		fmt.Sprintf(`Select the GPT model to use.
		Options: 0 (%s), 1 (%s)
		Default: 1`, agent.GPT3Dot5, agent.GPT4Dot5))
)

func init() {
	flag.Parse()
	if *ModelFlag == "0" {
		*ModelFlag = string(agent.GPT3Dot5)
	} else {
		*ModelFlag = string(agent.GPT4Dot5)
	}
}
