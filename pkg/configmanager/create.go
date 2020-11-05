// Package pathmanager manager the paths associated to the files wished to observed
package pathmanager

import (
	"fmt"
	"github.com/ctfrancia/go-dot/pkg/models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

var basic = `
# basic structure

repoURL: none
repoToken: none 
files:
	- type: neovim # manager type 
	  path: ~/.config/nvim/init.vim # path to the file # location of file
	  runCmd: nil # (optional) any run commands associated
	- type: zsh
	  path: ~/.zshrc
	  runCmd: nil
`

// CreateGodotDirPath creates the directory and config file where godot is stored
func CreateGodotDirPath(p string) error {
	err := os.MkdirAll(p, 0777)
	if err != nil {
		return err
	}

	fmt.Println("created config directory at: ", p)
	return nil
}

// ConfigCreate is for creating the congfig file that we will use later on the addition/deletion/update/etc
func ConfigCreate(cPath string) error { // right now just with zshfile

	f, err := os.Create(filepath.Join(cPath, "config.yaml"))
	if err != nil {
		return err
	}

	defer f.Close()

	fmt.Println("created config.yaml file")
	return nil
}

// CreateBasicConfig creates a boilerplate for how the config file looks
func CreateBasicConfig() error {
	// m := models.GoDotConfig{}
	usr, err := user.Current()
	if err != nil {
		return err
	}
	path := filepath.Join(usr.HomeDir, ".config", "config.yaml")
	err = ioutil.WriteFile(path, []byte(basic), 0777)
	if err != nil {
		return err
	}

	return nil
}
