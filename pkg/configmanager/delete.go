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

// DeleteConfigFolder deletes config folder
func DeleteConfigFolder() error {
	return nil
}
