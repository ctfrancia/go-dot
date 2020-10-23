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

	/*
		viper.SetConfigName("config")
		viper.SetConfigType("json")
		viper.AddConfigPath(configDirPath)

		// set default structure
		viper.SetDefault("godot", GodotDefault())

		err := viper.WriteConfig()
		if err != nil {
			log.Fatal(err)
		}

		viper.Set("godot.repoURL", "foo")
		viper.WriteConfig()
	*/

	/*
		err = viper.ReadInConfig() // Find and read the config file
		if err != nil {            // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	*/

	// TESTING PURPOSES OF JSON EDITING!!!
	// See()
	/*
		c := New()
		err = c.AddRepoURL("github.com/ctfrancia/.files")
		if err != nil {
			fmt.Println("error")
			log.Fatal(err)
		}
	*/
}

/*
this will be used for individual default locations of popular files
func osCheck() string {
	// check the user's OS and return the value

	return ""
}
*/
