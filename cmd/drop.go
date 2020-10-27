package cmd

import (
	// cm "github.com/ctfrancia/go-dot/pkg/configmanager"
	"fmt"
	"github.com/ctfrancia/go-dot/pkg/prompts"
	"github.com/spf13/cobra"
)

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "delete your config file and folder !!HOT!!",
	Long:  `drop works like you would do for a database, the fild and folder will be deleted and will need to be reinitialized`, Run: handleDropCmd,
}

func init() {
	rootCmd.AddCommand(dropCmd)
}

func handleDropCmd(cmd *cobra.Command, args []string) {
	dropC := prompts.DropConfig()
	if dropC {
		// _, err := cm.DeleteConfig()
		fmt.Println("yes was selected")
	}
	// TODO implement this cmd.
	// 1. Prompt user if they wish to continue,
	// 2. yes? remove file/folder
	// 3. ask if they want to restart
	// 4. if yes start the process over again on init new user
	// 5. if no then os.Exit(1) or something more appropriate
}
