package prompts

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// RestartPrompt asks if the user wishes to retart the config file and start from zero
func RestartPrompt() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("looks like you already have a config file with your paths, are you sure you want to start over? (y/n): ")
	text, _ := r.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	if text == "y" || text == "yes" {
		// call the pathmanager.
	}
}
