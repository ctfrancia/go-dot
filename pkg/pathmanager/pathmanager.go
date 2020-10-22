// Package pathmanager manager the paths associated to the files wished to observed
package pathmanager

import (
	// "fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

// userpaths defines how the paths look for each valid fiel
type Userpaths struct {
	zshPath  string
	nvimPath string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// New returns a new instance of userpaths to access it's methods
func New() *Userpaths {
	var up Userpaths
	return &up
}

// StartUp is this initial function that will check the user to see if the user is new or the root directory is already there
// if the user is new then it will start the prompt. If not then it will prepare the files
func StartUp() *Userpaths {
	var filepaths Userpaths

	exists := goDotFolderExists()
	if exists {
		return &filepaths
	}

	return &filepaths
}

func pathValidator(p, filename string) error {
	return nil
}

func goDotFolderExists() bool {
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

	return true
}

func (*Userpaths) AddPath(p string) error {

	return nil
}
