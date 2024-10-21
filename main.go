package main

import (
	"os"

	"github.com/AppleGamer22/cocainate/commands"
)

func main() {
	if err := commands.RootCommand.Execute(); err != nil {
		// _, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
