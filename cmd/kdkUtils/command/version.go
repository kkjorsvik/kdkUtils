package command

import (
	"flag"
	"fmt"
	"os"
)

var versionUsage = `Version.

Usage: kdkUtil version

`

var (
	build   = "???"
	version = "???"
)

var versionFunc = func(cmd *Command, args []string) {
	fmt.Printf("Version: v%s, Build: %s", version, build)
	os.Exit(0)
}

func VersionCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("version", flag.ExitOnError),
		Execute: versionFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, versionUsage)
	}

	return cmd
}
