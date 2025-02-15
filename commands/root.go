package commands

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/AppleGamer22/cocainate/session"
	"github.com/spf13/cobra"
)

var (
	duration time.Duration
	pid      int
	quiet    bool
)

var RootCommand = &cobra.Command{
	Use:     "cocainate",
	Short:   "keep screen awake",
	Long:    "keep screen awake",
	Version: Version,
	Args: func(cmd *cobra.Command, args []string) error {
		if pid != 0 && duration == 0 {
			return errors.New("process poling interval must be provided via the -d flag")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		s := session.New(duration, pid)

		if err := s.Start(); err != nil {
			return err
		}

		if err := s.Wait(quiet); err != nil {
			return err
		}

		if flag.Lookup("test.v") == nil {
			fmt.Print("\r")
		}

		return nil
	},
}

func init() {
	RootCommand.Flags().DurationVarP(&duration, "duration", "d", 0, "duration with units ns, us (or µs), ms, s, m, h")
	RootCommand.Flags().IntVarP(&pid, "pid", "p", 0, "a running process ID, duration (used as polling interval) must be provided")
	RootCommand.Flags().BoolVarP(&quiet, "quiet", "q", false, "hide progress bar")
	// RootCommand.MarkFlagsRequiredTogether("pid", "duration")
	RootCommand.SetVersionTemplate("{{.Version}}\n")
}
