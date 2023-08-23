package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kkjorsvik/kdkUtils/cmd/kdkUtils/command"
)

var usage = `Usage: kdkUtil command [options]

A simple tool to generate and manage custom templates

Options:

Commands:
  linux		Linux Related Stuff
  aws		Aws Related Stuff
  version	Prints version info to console
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	}

	flag.Parse()
	if flag.NArg() < 1 {
		usageAndExit("")
	}

	var cmd *command.Command

	switch os.Args[1] {
	case "linux":
		cmd = command.LinuxCommand()
	case "aws":
		cmd = command.AwsCommand()
	case "version":
		cmd = command.VersionCommand()
	default:
		usageAndExit(fmt.Sprintf("kdkUtil: '%s' is not a kdkUtil command.\n", os.Args[1]))
	}

	err := cmd.Init(os.Args[2:])
	if err != nil {
		fmt.Fprint(os.Stderr, "Error Init", "\n")
		os.Exit(1)
	}
	cmd.Run()
}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprint(os.Stderr, "\n")
	}

	flag.Usage()
	os.Exit(0)
}
