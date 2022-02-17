package cmd

import (
	"fmt"
	"os"
	"os/user"
	"text/tabwriter"

	"github.com/shirou/gopsutil/process"
	"github.com/spf13/cobra"
)

var all bool

var blacklistedProcesses = []string{
	"systemd",
	"sddm-helper",
	"pulseaudio",
	"xdg-document-portal",
	"gdm-session-worker",
	"gdm-x-session",
	"gnome-session-binary",
	"dbus-daemon",
	"at-spi-bus-launcher",
	"ibus-daemon",
	"gnome-shell",
	"ibus-memconf",
	"ibus-daemon",
	"ibus-extension-gtk3",
	"ibus-daemon",
	"gvfsd-trash",
	"gvfsd",
	"evolution-alarm-notify",
	"gnome-session-binary",
	"ibus-engine-simple",
	"ibus-daemon",
	"gsd-disk-utility-notify",
	"gnome-session-binary",
	"plasma-browser-integration-host",
	"nacl_helper",
}

var psCommand = &cobra.Command{
	Use:   "ps",
	Short: "list important processes",
	Long:  "list important processes",
	RunE: func(cmd *cobra.Command, args []string) error {
		u, err := user.Current()
		if err != nil {
			return err
		}

		processes, err := process.Processes()
		if err != nil {
			return err
		}

		writer := tabwriter.NewWriter(os.Stdout, 8, 8, 1, ' ', 0)
		defer writer.Flush()

		n, err := fmt.Fprintf(writer, "%s\t%s\t%s\n", "PID", "PPID", "CMD")
		if err != nil || n == 0 {
			return nil
		}

		processNames := make(map[string]bool)
		for _, p := range processes {
			pid := p.Pid
			curentPID := os.Getpid()
			if pid == int32(curentPID) {
				continue
			}

			parent, err := p.Parent()
			if err != nil {
				continue
			}
			parentName, err := parent.Name()
			if err != nil {
				continue
			}

			processName, err := p.Name()
			if err != nil {
				continue
			}

			var blacklisted bool
			for _, blacklistedProcess := range blacklistedProcesses {
				if processName == blacklistedProcess || parentName == blacklistedProcess {
					blacklisted = true
					break
				}
			}
			if blacklisted {
				continue
			}

			processUsername, err := p.Username()
			if err != nil || processUsername != u.Username {
				continue
			}

			if _, ok := processNames[processName]; all || !ok {
				processNames[processName] = true

				ppid, err := p.Ppid()
				if err != nil || ppid == 1 {
					continue
				}

				n, err := fmt.Fprintf(writer, "%d\t%d\t%s\n", pid, ppid, processName)
				if err != nil || n == 0 {
					return nil
				}
			}

		}

		return nil
	},
}

func init() {
	psCommand.Flags().BoolVarP(&all, "all", "a", false, "show all processes")

	RootCommand.AddCommand(psCommand)
}
