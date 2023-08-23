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
	if _, err := os.UserHomeDir(); err != nil {
		errAndExit("Failed to return user's home directory. Some reason this is important, maybe?")
	}
	fmt.Println("Hello from Linux Stuff")
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
