// Package pathmanager manager the paths associated to the files wished to observed
package pathmanager

import (
	"fmt"
	//	"github.com/ctfrancia/go-dot/pkg/models"
	//	"gopkg.in/yaml.v2"
	// "io/ioutil"
	"os"
	// "os/user"
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

// CreateGodot creates the directory and config file where godot is stored
func CreateGodot(p string) (int, error) {
	err := os.MkdirAll(p, 0777)
	if err != nil {
		return 0, err
	}

	f, err := os.Create(filepath.Join(p, "config.yaml"))
	if err != nil {
		return 0, err
	}

	b, err := f.Write([]byte(basic))
	if err != nil {
		return 0, err
	}

	fmt.Println("created config directory at: ", p)
	return b, nil
}
