// Package pathmanager manager the paths associated to the files wished to observed
package pathmanager

import (
	"fmt"
)

// userpaths defines how the paths look for each valid fiel
type userpaths struct {
	zshPath  string
	nvimPath string
}

// StartUp is this initial function that will check the user to see if the user is new or the root directory is already there
// if the user is new then it will start the prompt. If not then it will prepare the files
func StartUp() {
	fmt.Println("asdf")
}
