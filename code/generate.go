package main

import (
	"fmt"
	"os"
	"strings"
)

// 1. SplitInput — only responsible for splitting input into segments
func SplitInputs(str string) []string {
	return strings.Split(str, "\\n")
}

// 2. renderLines — only responsible for rendering ONE segment into 8 rows
func renderLines(segment string, blocksMap map[rune][]string) []string {
	var output []string
	for row := 0; row < 8; row++ {
		var result strings.Builder
		for _, ch := range segment {
			result.WriteString(blocksMap[ch][row])
		}
		output = append(output, result.String())
	}
	return output
}

// 3. Generate — orchestrates both, handles empty segments, builds final string
func Generate(input string, banner map[rune][]string) string {
	var output strings.Builder
	segments := SplitInput(input)

	for i, segment := range segments {
		if segment == "" {
			if i < len(segments)-1 {
				output.WriteString("\n")
			}
			continue
		}
		for _, ch := range segment {
			if _, ok := banner[ch]; !ok {
				fmt.Fprintf(os.Stderr, "Non ASCII character %q\n", ch)
				os.Exit(1)
			}
		}
		rows := renderLines(segment, banner)
		for _, row := range rows {
			output.WriteString(row + "\n")
		}
	}
	return output.String()
}
