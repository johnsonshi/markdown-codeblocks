package parser

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

type CodeBlock struct {
	// The language of the code block.
	Language string
	// The code block's literal content.
	Literal string
}

// ParseMarkdownCodeBlocks parses all fenced code blocks within a Markdown file.
// A language filter can be specified to only parse code blocks that match the language filter.
func ParseMarkdownCodeBlocks(md []byte, languageFilter string) []CodeBlock {
	codeBlocks := []CodeBlock{}

	p := parser.New()
	doc := p.Parse(md)

	ast.WalkFunc(doc, func(node ast.Node, entering bool) ast.WalkStatus {
		if codeBlock, ok := node.(*ast.CodeBlock); ok {
			if languageFilter == "" || string(codeBlock.Info) == languageFilter {
				codeBlocks = append(codeBlocks, CodeBlock{
					Language: string(codeBlock.Info),
					Literal:  string(codeBlock.Literal),
				})
			}
		}
		return ast.GoToNext
	})

	return codeBlocks
}
