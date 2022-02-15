package cmd

import (
	"github.com/spf13/cobra"
)

var psCommand = &cobra.Command{
	Use:   "ps",
	Short: "list important processes",
	Long:  "list important processes",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	RootCommand.AddCommand(psCommand)
}
