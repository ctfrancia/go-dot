// Package pathmanager manager the paths associated to the files wished to observed
package pathmanager

import (
	"fmt"
	"os"
	"path/filepath"
)

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
