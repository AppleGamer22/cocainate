package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hours, minutes, seconds uint

var RootCommand = &cobra.Command{
	Use:   "cocainate",
	Short: "keep screen awake",
	Long:  "keep screen awake",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("hi")
		return nil
	},
}

func init() {
	RootCommand.Flags().UintVarP(&hours, "hours", "h", 0, "number of hours")
	RootCommand.Flags().UintVarP(&minutes, "minutes", "m", 0, "number of minutes")
	RootCommand.Flags().UintVarP(&seconds, "seconds", "s", 0, "seconds")
}
