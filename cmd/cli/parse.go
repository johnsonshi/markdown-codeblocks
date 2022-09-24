package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/johnsonshi/markdown-codeblocks/parser"
)

type parseCmdOpts struct {
	stdin          io.Reader
	stdout         io.Writer
	stderr         io.Writer
	file           string
	languageFilter string
	execute        bool
	interactive    bool
}

func newParseCmd(stdin io.Reader, stdout io.Writer, stderr io.Writer, args []string) *cobra.Command {
	opts := &parseCmdOpts{
		stdin:  stdin,
		stdout: stdout,
		stderr: stderr,
	}

	cobraCmd := &cobra.Command{
		Use:   "parse",
		Short: "Parses and optionally executes fenced code blocks within a Markdown file. A code block language filter can be specified.",
		RunE: func(_ *cobra.Command, args []string) error {
			return opts.run()
		},
	}

	f := cobraCmd.Flags()

	f.StringVarP(&opts.file, "file", "f", "", "path to input Markdown file")
	cobraCmd.MarkFlagRequired("file")

	f.StringVarP(&opts.languageFilter, "language-filter", "l", "", "only code blocks that match this language will be parsed")

	f.BoolVarP(&opts.execute, "execute", "e", false, "execute the parsed code blocks")
	f.BoolVarP(&opts.interactive, "interactive", "i", false, "interactive mode when executing the parsed code blocks (requires -e)")

	return cobraCmd
}

func (opts *parseCmdOpts) run() error {
	md, err := os.ReadFile(opts.file)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(opts.stdin)

	codeBlocks := parser.ParseMarkdownCodeBlocks(md, opts.languageFilter)

	for _, codeBlock := range codeBlocks {
		if opts.execute && codeBlock.Language == opts.languageFilter {
			fmt.Fprintf(opts.stdout, "[*] Executing code block with language: %s\n", codeBlock.Language)
			fmt.Fprintf(opts.stdout, "[*] Code block:\n%s\n", codeBlock.Literal)

			if opts.interactive {
				fmt.Fprintf(opts.stdout, "[*] Press enter to execute the code block...")
				scanStatus := scanner.Scan()
				if err = scanner.Err(); err != nil {
					return err
				}
				// Scanner hit EOF.
				if !scanStatus {
					return nil
				}
			}

			// Execute code block by calling "<language>" with the code block.
			cmd := exec.Command(codeBlock.Language, "-c", codeBlock.Literal)
			cmd.Stdin = opts.stdin
			cmd.Stdout = opts.stdout
			cmd.Stderr = opts.stderr

			fmt.Fprint(opts.stdout, "[*] Output:\n")
			err = cmd.Run()
			if err != nil {
				return err
			}

			fmt.Fprint(opts.stdout, "\n==============================================\n")
		} else {
			opts.stdout.Write([]byte(codeBlock.Literal))
		}
	}

	return nil
}
