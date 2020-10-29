// Package pathmanager manager the paths associated to the files wished to observed
package pathmanager

import (
	"fmt"
	// "github.com/spf13/viper"
	// "log"
	"os"
	"path/filepath"
)

// CreateGodotDirPath creates the directory and config file where godot is stored
func CreateGodotDirPath(p string) error {
	err := os.MkdirAll(p, 0777)
	if err != nil {
		return err
	}

	fmt.Println("created config directory")
	return nil
}

// ConfigCreate is for creating the congfig file that we will use later on the addition/deletion/update/etc
func ConfigCreate(cPath string) error { // right now just with zshfile

	f, err := os.Create(filepath.Join(cPath, "config.json"))
	if err != nil {
		return err
	}

	defer f.Close()

	fmt.Println("created config.json file")
	return nil
}

/*
this will be used for individual default locations of popular files
func osCheck() string {
	// check the user's OS and return the value

	return ""
}
*/
