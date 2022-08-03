package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
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
		file, err := os.OpenFile(path, os.O_RDWR, 0)
		if err != nil {
			return err
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		var buffer bytes.Buffer
		for scanner.Scan() {
			line := scanner.Text()
			updatedLine := strings.ReplaceAll(line, before, after)
			if _, err := buffer.WriteString(updatedLine); err != nil {
				fmt.Println(err)
			}
			if _, err := buffer.WriteRune('\n'); err != nil {
				fmt.Println(err)
			}
		}
		if err := scanner.Err(); err != nil {
			return err
		}
		file.Seek(0, io.SeekStart)
		_, err = io.Copy(file, &buffer)
		return err
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
