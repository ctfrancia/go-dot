package cmd

import (
	"github.com/ctfrancia/go-dot/pkg/pathmanager"
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
	Run:   handleInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleInit(cmd *cobra.Command, args []string) {

	usr, err := user.Current()
	check(err)

	godotDir := filepath.Join(usr.HomeDir, ".config", "godot")
	zshDefaultPathUnix := filepath.Join(usr.HomeDir, "/.zshrc")

	if !goDotConfigExists(godotDir) {
		pathmanager.CreateGodotDirPath(godotDir)
		pathmanager.ConfigCreate(godotDir, zshDefaultPathUnix)
		// prompts.initNewUser()
		return
	}
	// the file exists
	// prompt user if they wish to start over again
	// promptToStartOverAgain()
	// if userWantsToStartOver() {}
}

func goDotConfigExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
