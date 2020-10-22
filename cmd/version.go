package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Godot",
	Long:  `All software has versions. This is Godot's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		fmt.Println("Godot A Dot file Manager v0.0.1 -- HEAD")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
