package command

import (
	"flag"
	"fmt"
	"os"
)

var linuxUsage = `Linux Stuff.

Usage: kdkUtil linux [OPTIONS] command

Options:
	-h Help
`

var linuxFunc = func(cmd *Command, args []string) {
	// Unsure if I need this yet
}

func LinuxCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("linux", flag.ExitOnError),
		Execute: linuxFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, linuxUsage)
	}

	return cmd
}
