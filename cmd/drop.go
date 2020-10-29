package cmd

import (
	"fmt"
	cm "github.com/ctfrancia/go-dot/pkg/configmanager"
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
		err := cm.DeleteConfig()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("config deleted successfully")
	}
}
