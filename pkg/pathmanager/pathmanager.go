// Package pathmanager manager the paths associated to the files wished to observed
package pathmanager

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	// "os/user"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateGodotDirPath(p string) {
	err := os.MkdirAll(p, 0777)
	check(err)

	f, err := os.Create(filepath.Join(p, "config.json"))
	check(err)
	defer f.Close()

	fmt.Println("in init true statement")
}

// ConfigCreate is for creating the congfig file that we will use later on the addition/deletion/update/etc
func ConfigCreate(configDirPath string, zPath string) { // right now just with zshfile
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configDirPath)

	// set default paths
	// viper.SetDefault("zshPath", zPath)
	viper.SetDefault("config", GodotDefault())

	viper.WriteConfig()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func osCheck() string {
	// check the user's OS and return the value

	return ""
}
