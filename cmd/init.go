package cmd

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new go-dot folder",
	Long:  `Will create a new go-dot folder as though it is your first time with the application`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO this here invokes the func where we start the prompt question again
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
