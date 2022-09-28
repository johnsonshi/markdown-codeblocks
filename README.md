# Markdown Code Blocks

Parses and optionally executes all [fenced code blocks](https://www.markdownguide.org/extended-syntax/#fenced-code-blocks) within a Markdown file.

A language filter can be applied to only parse and execute fenced code blocks from a specific language.
This requires the fenced code blocks to have the language specifier after the triple backticks (`` ```<language> ``).

[![asciicast](https://asciinema.org/a/1so5LtPQK95HrAiijtHKYyI4I.svg)](https://asciinema.org/a/1so5LtPQK95HrAiijtHKYyI4I)

## Usage

```
Usage:
  markdown-codeblocks parse [flags]

Flags:
  -e, --execute                  execute the parsed code blocks
  -f, --file string              path to input Markdown file
  -h, --help                     help for parse
  -i, --interactive              interactive mode when executing the parsed code blocks (requires -e)
  -l, --language-filter string   only code blocks that match this language will be parsed
```

## Examples

To **parse** all `bash` fenced code blocks:

```bash
bin/markdown-codeblocks parse \
        --file markdown.md \
        --language-filter bash
```

To **parse and execute** all `bash` fenced code blocks:

```bash
bin/markdown-codeblocks \
    parse \
    -f parser/testdata/multiple_code_blocks_with_language_filter_that_exists_in_some_code_blocks.md \
    --language-filter bash \
    --execute
```

To **parse and execute** all `bash` fenced code blocks **in an interactive session (prompting before executing each block)**:

```bash
bin/markdown-codeblocks \
    parse \
    -f parser/testdata/multiple_code_blocks_with_language_filter_that_exists_in_some_code_blocks.md \
    --language-filter bash \
    --execute \
    --interactive
```
