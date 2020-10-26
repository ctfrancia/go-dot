package cmd

import (
	"github.com/spf13/cobra"
	/*
		"fmt"
		cm "github.com/ctfrancia/go-dot/pkg/configmanager"
		"log"
		"os"
		"os/user"
		"path/filepath"
	*/)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a file to be watched",
	Long:  `adds a file to the config file, after the flag pass in the path to the file`,
	Run:   handleAddCmd,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func handleAddCmd(cmd *cobra.Command, args []string) {
	//TODO handle this command
}
