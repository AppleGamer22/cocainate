package cmd

import (
	"time"

	"github.com/AppleGamer22/cocainate/internal"
	"github.com/spf13/cobra"
)

var duration time.Duration
var pid int

var RootCommand = &cobra.Command{
	Use:   "cocainate",
	Short: "keep screen awake",
	Long:  "keep screen awake",
	RunE: func(cmd *cobra.Command, args []string) error {
		session := internal.Session{
			PID:      pid,
			Duration: duration,
		}
		if err := session.Start(); err != nil {
			return err
		}
		return session.Wait()
	},
}

func init() {
	RootCommand.Flags().DurationVarP(&duration, "duration", "d", 0, "duration with units ns, us (or µs), ms, s, m, h")
	RootCommand.Flags().IntVar(&pid, "pid", 0, "process ID")
}
