package prompts

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// DropConfig prompts the user if they want to delete the godot folder and config
func DropConfig() bool {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Warning: you are about to delete the godot directory and config file, are you sure you want to continue? (y/n): ")

	text, _ := r.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	if text == "y" || text == "yes" {
		return true
	}

	return false
}
