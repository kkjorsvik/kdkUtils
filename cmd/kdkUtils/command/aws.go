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
	homeDir, err := os.UserHomeDir()
	if err != nil {
		errAndExit("Failed to return user's home directory.")
	}
	fmt.Printf("Home Directory: %s", homeDir)
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
