package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var duration time.Duration

var RootCommand = &cobra.Command{
	Use:   "cocainate",
	Short: "keep screen awake",
	Long:  "keep screen awake",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("hi")
		fmt.Printf("%v\n", duration)
		return nil
	},
}

func init() {
	RootCommand.Flags().DurationVarP(&duration, "duration", "d", 0, "duration with units ns, us (or µs), ms, s, m, h")
}
