package cmd

import (
	"fmt"
	cm "github.com/ctfrancia/go-dot/pkg/configmanager"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new go-dot folder",
	Long:  `Will create a new go-dot folder as though it is your first time with the application`,
	Run:   handleInitCmd,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleInitCmd(cmd *cobra.Command, args []string) {

	usr, err := user.Current()
	check(err)

	godotDir := filepath.Join(usr.HomeDir, ".config", "godot")

	if goDotConfigExists(godotDir) {
		fmt.Println("file exists")
		// the file exists
		// prompt user if they wish to start over again
		// promptToStartOverAgain()
		// if userWantsToStartOver() {}
		return
	}

	fmt.Println("Thank you for using go-dot, one second while we create your config path")
	err = cm.CreateGodotDirPath(godotDir)
	check(err)

	err = cm.ConfigCreate(godotDir)
	check(err)
	// prompts.initNewUser()
}

func goDotConfigExists(p string) bool {

	if _, err := os.Stat(p); !os.IsNotExist(err) {
		// file exists
		return true
	}
	return false
}
