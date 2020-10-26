package pathmanager

import (
	"encoding/json"
	"fmt"
	m "github.com/ctfrancia/go-dot/pkg/models"
	"io/ioutil"
	"log"
)

var gdConfig m.GdotConfig
var gdotC m.GDotC

// New returns the pointer to the struct for methods
func New() *gdConfig {
	return &gdConfig
}

// AddRepoURL adds the url to where the dotfiles are being remotely stored by GitHub/BitBucket/Gitlab/etc.
func (c *GdotConfig) AddRepoURL(p string) error {
	content, err := ioutil.ReadFile("/Users/christian.francia/.config/godot/config.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &gdotC)
	if err != nil {
		return err
	}

	gdotC.Godot.RepoURL = p

	fmt.Printf("after unmarshal %#v", gdotC)

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

	err = json.Unmarshal(content, &gdotC)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", gdotC)

	// fmt.Println("---->", gdotC.Godot.RepoURL)

}

// AddToConfig takes the key and the value that needs to be added
func (c *GdotConfig) AddToConfig(k, v string) error {
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
