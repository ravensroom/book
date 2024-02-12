package flag

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	GitFlag = flag.String("g", "", "Git show/diff command\nExample: -g \"git show\" or -g \"git diff\"")
)

var ALLOWED_GIT_COMMANDS = []string{"git show", "git diff"}

func init() {
	flag.Parse()

	if *GitFlag != "" {
		parts := strings.Split(*GitFlag, " ")
		if len(parts) < 2 {
			fmt.Println("Please wrap the git command in quotes and include both git and its subcommand.\nExample: -g \"git show\" or -g \"git diff --cached\"")
			os.Exit(1)
		}
		command := strings.Join(parts[:2], " ")

		found := false
		for _, allowedCommand := range ALLOWED_GIT_COMMANDS {
			if command == allowedCommand {
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Error: Unsupported Git command. Please use one of the following commands:", ALLOWED_GIT_COMMANDS)
			os.Exit(1)
		}
	} else {
		fmt.Println("Git command not fonud. Default to git show")
		*GitFlag = "git show"
	}
}
