package main

import (
	/*
		"fmt"
		// "github.com/google/go-github/v32/github"
		// "flag"
		"bufio"
		// "io/ioutil"
		"log"
		"os"
		"os/user"
		"path/filepath"
		"strings"
	*/
	"github.com/ctfrancia/go-dot/pkg/pathmanager"
)

type dotFiles struct {
	zshrcPath string
}

func main() {
	// var df dotFiles
	// var zshrcPath string
	// var initVimPath string
	pathmanager.StartUp()

	/*
		usr, err := user.Current()
		check(err)

		godotFile := filepath.Join(usr.HomeDir, "/godot")
		check(err)

		if _, err := os.Stat(godotFile); os.IsNotExist(err) {
			err = os.MkdirAll(godotFile, 0777)
			check(err)

			f, err := os.Create(filepath.Join(godotFile, "config.json"))
			check(err)
			defer f.Close()
		}
	*/

	/*
		r := bufio.NewReader(os.Stdin)

		fmt.Print("do you wish to track your zshrc file(y/n)?: ")
		text, _ := r.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		if text == "y" || text == "yes" {
			fmt.Print("is it located in the default location(y/n)?: ")
			text, _ = r.ReadString('\n')
			text = strings.TrimSuffix(text, "\n")
			if text == "y" || text == "yes" {
				df.zshrcPath = filepath.Join(usr.HomeDir, ".zshrc")
				fmt.Println(df.zshrcPath)
			}
		}
	*/
}
