package main

import (
	"fmt"
	"os"

	"github.com/AppleGamer22/cocainate/cmd"
)

func main() {
	if err := cmd.RootCommand.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
