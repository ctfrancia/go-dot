package pathmanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

var gdConfig GdotConfig
var gdotC GDotC

// New returns the pointer to the struct for methods
func New() *GdotConfig {
	return &gdConfig
}

// AddRepoURL adds the url to where the dotfiles are being remotely stored by GitHub/BitBucket/etc.
func (c *GdotConfig) AddRepoURL(p string) error {
	content, err := ioutil.ReadFile("/Users/christian.francia/.config/godot/config.json")
	if err != nil {
		// log.Fatal(err)
		return err
	}

	err = json.Unmarshal(content, &gdotC)
	if err != nil {
		// log.Fatal(err)
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

func save() {
}
