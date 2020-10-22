package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "godot",
	Short: "Godot is a useful tool for managing your dot files",
	Long: `Godot is cli managment tool that is used for pull and pushing your dot files
			spread throughout your computer to centralize and fix the laborious task of manually
			copy, pasting, and saving your dot files`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do stuff here
	},
}

// Execute does something still figuring out what, following the docs
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
