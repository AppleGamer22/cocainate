package commands

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	Version = "development"
	Hash    = "development"
	Date    = "2006-01-02"

	verbose bool

	versionCommand = &cobra.Command{
		Use:   "version",
		Short: "print version",
		Long:  "print version",
		Run: func(cmd *cobra.Command, args []string) {
			if verbose {
				if Version != "development" {
					fmt.Printf("version: \t%s\n", Version)
				}

				if Hash != "development" {
					fmt.Printf("commit: \t%s\n", Hash)
				}
				fmt.Printf("compiler: \t%s (%s)\n", runtime.Version(), runtime.Compiler)
				fmt.Printf("platform: \t%s/%s\n", runtime.GOOS, runtime.GOARCH)
			} else {
				fmt.Println(Version)
			}
		},
	}
)

func init() {
	versionCommand.Flags().BoolVarP(&verbose, "verbose", "v", false, "version, git commit hash, compiler version & platform")
	RootCommand.AddCommand(versionCommand)
}
