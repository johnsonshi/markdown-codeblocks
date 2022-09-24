package main

import (
	"flag"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd(stdin io.Reader, stdout io.Writer, stderr io.Writer, args []string) *cobra.Command {
	cobraCmd := &cobra.Command{
		Use:   "markdown-cb",
		Short: "Parses and optionally executes fenced code blocks within a Markdown file.",
	}

	cobraCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	flags := cobraCmd.PersistentFlags()

	cobraCmd.AddCommand(
		newParseCmd(stdin, stdout, stderr, args),
	)

	_ = flags.Parse(args)

	return cobraCmd
}

func execute() {
	rootCmd := newRootCmd(os.Stdin, os.Stdout, os.Stderr, os.Args[1:])
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
