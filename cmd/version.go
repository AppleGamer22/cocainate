package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "1.0.0"

var verbose bool

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "print version",
	Long:  "print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}

func init() {
	versionCommand.Flags().BoolVarP(&verbose, "verbose", "v", false, "-v or --verbose")

	RootCommand.AddCommand(versionCommand)
}
