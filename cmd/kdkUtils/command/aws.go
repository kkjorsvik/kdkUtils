package command

import (
	"flag"
	"fmt"
	"os"
)

var awsUsage = `AWS Stuff.

Usage: kdkUtil aws [OPTIONS] command

Options:
	-h Help
`

var awsFunc = func(cmd *Command, args []string) {
	if _, err := os.UserHomeDir(); err != nil {
		errAndExit("Failed to return user's home directory.")
	}
	fmt.Println("Hello from AWS Stuff!")
}

func AwsCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("aws", flag.ExitOnError),
		Execute: awsFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, awsUsage)
	}

	return cmd
}
