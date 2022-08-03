package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var before, after string

var rootCommand = cobra.Command{
	Args: func(_ *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("the file path argument is required")
		}
		_, err := os.Stat(args[0])
		return err
	},
	PreRunE: func(_ *cobra.Command, args []string) error {
		if before == "" || (before == "" && after == "") {
			return errors.New("invalid replacement argument(s)")
		}
		return nil
	},
	RunE: func(_ *cobra.Command, args []string) error {
		path := args[0]
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			updatedLine := strings.ReplaceAll(line, before, after)
			fmt.Println(updatedLine)
		}
		return nil
	},
}

func init() {
	rootCommand.Flags().StringVarP(&before, "before", "b", "", "")
	rootCommand.Flags().StringVarP(&after, "after", "a", "", "")
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
