package parser

import (
	"encoding/json"
	"flag"
	"os"
	"testing"
)

var (
	update = flag.Bool("update", false, "update the golden files of this test")
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func Test_ParseMarkdownCodeBlocks(t *testing.T) {
	var tests = []struct {
		name           string
		languageFilter string
	}{
		{
			"empty",
			"",
		},
		{
			"no_code_blocks_with_no_language_filter",
			"",
		},
		{
			"no_code_blocks_with_language_filter",
			"go",
		},
		{
			"single_code_block_with_no_language_filter",
			"",
		},
		{
			"single_code_block_with_language_filter_that_exists_in_the_code_block",
			"bash",
		},
		{
			"single_code_block_with_language_filter_that_does_not_exist_in_the_code_block",
			"fortran",
		},
		{
			"multiple_code_blocks_with_no_language_filter",
			"",
		},
		{
			"multiple_code_blocks_with_language_filter_that_exists_in_some_code_blocks",
			"bash",
		},
		{
			"multiple_code_blocks_with_language_filter_that_does_not_exist_in_any_code_blocks",
			"fortran",
		},
	}

	for _, test := range tests {
		inputMd, err := os.ReadFile("testdata/" + test.name + ".md")
		if err != nil {
			t.Fatal(err)
		}
		codeBlocks := ParseMarkdownCodeBlocks(inputMd, test.languageFilter)
		codeBlocksJson, err := json.MarshalIndent(codeBlocks, "", "  ")
		if err != nil {
			t.Fatal(err)
		}

		expected := goldenValue(t, test.name, string(codeBlocksJson), *update)

		if string(codeBlocksJson) != expected {
			t.Errorf("Expected:\n%s\nActual:\n%s", expected, codeBlocksJson)
		}
	}
}

func goldenValue(t *testing.T, goldenFile string, actual string, update bool) string {
	t.Helper()
	goldenPath := "testdata/" + goldenFile + ".golden"

	if update {
		err := os.WriteFile(goldenPath, []byte(actual), 0644)
		if err != nil {
			t.Fatalf("Error writing to file %s: %s", goldenPath, err)
		}

		return actual
	}

	content, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Error opening file %s: %s", goldenPath, err)
	}
	return string(content)
}
