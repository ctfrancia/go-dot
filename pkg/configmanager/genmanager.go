// Package pathmanager manager the paths associated to the files wished to observed
package pathmanager

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// CreateGodotDirPath creates the directory and config file where godot is stored
func CreateGodotDirPath(p string) {
	err := os.MkdirAll(p, 0777)
	check(err)

	f, err := os.Create(filepath.Join(p, "config.json"))
	check(err)
	defer f.Close()

	fmt.Println("created config file")
}

// ConfigCreate is for creating the congfig file that we will use later on the addition/deletion/update/etc
func ConfigCreate(configDirPath string) { // right now just with zshfile
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configDirPath)

	// set default structure
	viper.SetDefault("godot", GodotDefault())

	err := viper.WriteConfig()
	if err != nil {
		log.Fatal(err)
	}

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
func osCheck() string {
	// check the user's OS and return the value

	return ""
}
*/
