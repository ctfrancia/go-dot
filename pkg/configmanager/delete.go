package pathmanager

import (
	"os"
	"os/user"
	"path/filepath"
)

// DeleteConfigFile deletes the config.json file
func DeleteConfigFile() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}

	err = os.Remove(filepath.Join(usr.HomeDir, ".config", "godot", "config.json"))
	if err != nil {
		return err
	}

	return nil
}

// DeleteConfig deletes config folder
func DeleteConfig() error {

	usr, err := user.Current()
	if err != nil {
		return err
	}

	err = os.Remove(filepath.Join(usr.HomeDir, ".config"))
	if err != nil {
		return err
	}

	return nil
}
