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

		writer := tabwriter.NewWriter(os.Stdout, 8, 8, 0, ' ', 0)
		defer writer.Flush()

		n, err := fmt.Fprintf(writer, "%s\t%s\t%s\n", "PID", "PPID", "CMD")
		if err != nil || n == 0 {
			return nil
		}

		processNames := make(map[string]bool)
		for _, p := range processes {
			processUsername, err := p.Username()
			if err != nil || processUsername != u.Username {
				continue
			}

			processName, err := p.Name()
			if err != nil {
				continue
			}

			if _, ok := processNames[processName]; all || !ok {
				processNames[processName] = true

				pid := p.Pid
				ppid, err := p.Ppid()
				if err != nil {
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
