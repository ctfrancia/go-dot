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

// TODO move this to it's own error handling pkg, because this doesn't provide any further detail
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
		fmt.Println("if you want to start over type: 'godot restart' or 'godot modify <file eg: 'zshrc'>' or 'godot -h'")
		return
	}

	fmt.Println("Thank you for using go-dot, one second while we create your config path")
	_, err = cm.CreateGodot(godotDir)
	check(err)

	fmt.Println("initialization of the config has completed, please use '-h' to see flags to modify your config file")
}

func goDotConfigExists(p string) bool {
	if _, err := os.Stat(p); !os.IsNotExist(err) {
		// file exists
		return true
	}
	return false
}
