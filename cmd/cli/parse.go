package main

import (
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
	inputFilePath  string
	languageFilter string
	execute        bool
}

func newParseCmd(stdin io.Reader, stdout io.Writer, stderr io.Writer, args []string) *cobra.Command {
	opts := &parseCmdOpts{
		stdin:  stdin,
		stdout: stdout,
		stderr: stderr,
	}

	cobraCmd := &cobra.Command{
		Use:   "parse",
		Short: "Parses and concatenates all fenced code blocks within a Markdown file. A code block language filter can be specified.",
		RunE: func(_ *cobra.Command, args []string) error {
			return opts.run()
		},
	}

	f := cobraCmd.Flags()

	f.StringVarP(&opts.inputFilePath, "input", "i", "", "input file path")
	cobraCmd.MarkFlagRequired("input")

	f.StringVarP(&opts.languageFilter, "language-filter", "l", "", "only code blocks that match this language will be parsed")

	f.BoolVarP(&opts.execute, "execute", "e", false, "execute the parsed code blocks")

	return cobraCmd
}

func (opts *parseCmdOpts) run() error {
	md, err := os.ReadFile(opts.inputFilePath)
	if err != nil {
		return err
	}

	codeBlocks := parser.ParseMarkdownCodeBlocks(md, opts.languageFilter)

	for _, codeBlock := range codeBlocks {
		if opts.execute && codeBlock.Language == opts.languageFilter {
			fmt.Fprintf(opts.stdout, "[*] Executing code block with language: %s\n", codeBlock.Language)
			fmt.Fprintf(opts.stdout, "[*] Code block:\n%s\n", codeBlock.Literal)

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

			fmt.Fprint(opts.stdout, "==============================================\n")
		} else {
			opts.stdout.Write([]byte(codeBlock.Literal))
		}
	}

	return nil
}
