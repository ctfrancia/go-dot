package pathmanager

import (
	"encoding/json"
	"fmt"
	m "github.com/ctfrancia/go-dot/pkg/models"
	"io/ioutil"
	"log"
)

// ConfigModel represents our config structure
type ConfigModel struct {
	CM *m.GoDotConfig
}

var config ConfigModel

// var Config m.GdotConfig

// New returns the pointer to the struct for methods
func New() *ConfigModel {
	return &ConfigModel{}
}

// AddRepoURL adds the url to where the dotfiles are being remotely stored by GitHub/BitBucket/Gitlab/etc.
func (c *ConfigModel) AddRepoURL(p string) error {
	// TODO refactor this so we have more SOC
	content, err := ioutil.ReadFile("/Users/christian.francia/.config/godot/config.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		return err
	}

	fmt.Printf("%v", content)
	// refactor to interact with the manager, returns *GoDotConfig
	// gdotC.Godot.RepoURL = p

	// fmt.Printf("after unmarshal %#v", gdotC)

	return nil
}

// See returns the json for viewing
func See() {
	content, err := ioutil.ReadFile("/Users/christian.francia/.config/godot/config.json")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(content))

	fmt.Println("before unmarshal", string(content))

	/*
		err = json.Unmarshal(content, &gdotC)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", gdotC)

		// fmt.Println("---->", gdotC.Godot.RepoURL)
	*/

}

// AddToConfig takes the key and the value that needs to be added
func (c *ConfigModel) AddToConfig(k, v string) error {
	return nil
}

// Save will update your config
func Save() error {
	return nil
}

// ModifyConfig takes the type eg: "Vim" and will update the
func ModifyConfig() error {
	return nil
}
