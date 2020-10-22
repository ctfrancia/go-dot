package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "godot",
	Short: "Godot is a useful tool for managing your dot files",
	Long: `Godot is cli managment tool that is used for pull and pushing your dot files
			spread throughout your computer to centralize and fix the laborious task of manually
			copy, pasting, and saving your dot files`,
	Run: handleCmd,
}

// Execute calls the root level of cobra for initialization
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func handleCmd(cmd *cobra.Command, args []string) {
	r := bufio.NewReader(os.Stdin)
	// Do stuff here
	fmt.Print("Thanks for using Godot, is there anything I can help you with? Type a flag or run 'godot -h' to see a list of available commands")
	text, _ := r.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
}
