package main

import (
	"os"

	"github.com/AppleGamer22/cocainate/cmd"
)

func main() {
	if err := cmd.RootCommand.Execute(); err != nil {
		// _, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
