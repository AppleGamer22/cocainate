package cmd

import (
	"time"

	"github.com/AppleGamer22/cocainate/session"
	"github.com/spf13/cobra"
)

var duration time.Duration
var pid int

var RootCommand = &cobra.Command{
	Use:   "cocainate",
	Short: "keep screen awake",
	Long:  "keep screen awake",
	RunE: func(cmd *cobra.Command, args []string) error {
		s := session.New(pid, duration)

		if err := s.Start(); err != nil {
			return err
		}
		return s.Wait()
	},
}

func init() {
	RootCommand.Flags().DurationVarP(&duration, "duration", "d", 0, "duration with units ns, us (or µs), ms, s, m, h")
	// RootCommand.Flags().IntVar(&pid, "pid", 0, "process ID")
}
