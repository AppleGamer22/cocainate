package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var Version = "development"
var Hash = "development"

var verbose bool

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "print version",
	Long:  "print version",
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Printf("version: \t%s\n", Version)
			fmt.Printf("commit: \t%s\n", Hash)
			fmt.Printf("compiler: \t%s (%s)\n", runtime.Version(), runtime.Compiler)
			fmt.Printf("platform: \t%s/%s\n", runtime.GOARCH, runtime.GOOS)
		} else {
			fmt.Println(Version)
		}
	},
}

func init() {
	versionCommand.Flags().BoolVarP(&verbose, "verbose", "v", false, "-v or --verbose")

	RootCommand.AddCommand(versionCommand)
}
