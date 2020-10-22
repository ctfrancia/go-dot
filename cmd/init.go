package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
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
	// fmt.Print("Thanks for using Godot, is there anything I can help you with? Type a flag or run 'godot -h' to see a list of available commands")

	usr, err := user.Current()
	check(err)

	godotFile := filepath.Join(usr.HomeDir, "/godot")
	check(err)

	if goDotConfigExists(godotFile) == false {
		createGoDotPath(godotFile)
		createConfigFile()
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

func createGoDotPath(p string) {
	err := os.MkdirAll(p, 0777)
	check(err)

	f, err := os.Create(filepath.Join(p, "config.json"))
	check(err)
	defer f.Close()

	fmt.Println("in init true statement")
}

func createConfigFile() {
	r := bufio.NewReader(os.Stdin)
	text, _ := r.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	// File exists, prompt user if they want to overwrite file
	// if they do, delete file and create again
	// github.com/spf13/viper for config creation
}
