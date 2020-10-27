package prompts

import (
	"bufio"
	"fmt"
	"github.com/ctfrancia/go-dot/pkg/configmanager"
	"os"
	// "path/filepath"
	"strings"
)

var filesToTrack []string = []string{"zsh", "init.vim"}

// StartQuestions function will initialize the CLI questions to fetch the paths from os.Stdin
func StartQuestions() {
	pm := pathmanager.New()
	fmt.Println(pm)
	r := bufio.NewReader(os.Stdin)

	/*
		for i, f := range filesToTrack {
		}
	*/

	// TODO this questionnaire needs to be moved into the loop above
	fmt.Print("do you wish to track your zshrc file(y/n)?: ")
	text, _ := r.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	if text == "y" || text == "yes" {
		fmt.Print("is it located in the default location(y/n)?: ")
		text, _ = r.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		if text == "y" || text == "yes" {
			/*
				df.zshrcPath = filepath.Join(usr.HomeDir, ".zshrc")
				fmt.Println(df.zshrcPath)
			*/
		}
	}
}

/*
func defaultPathLocation(pm *Userpaths) {
	pm.zshrcPath = filepath.Join(usr.HomeDir, ".zshrc")
	fmt.Println(pm.zshrcPath)
}
*/
