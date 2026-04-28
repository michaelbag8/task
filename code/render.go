package main

import (
	"fmt"
	"strings"
)
func renderBanner(input string, blocks map[rune][]string) []string {
	var output []string
	segments := strings.Split(input, "\\n")
	for i, segment := range segments {
		if segment == "" {
			if i < len(segments)-1 {
				fmt.Println()
			}
			continue
		}

		for _, ch := range segment {
			if _, exist := blocks[ch]; !exist {
				fmt.Println("non ascii character")
			}
		}
		for row := 0; row < 8; row++ {
			var result strings.Builder
			for _, ch := range segment {
				result.WriteString(blocks[ch][row])
			}

			output = append(output, result.String())

			
		}
		

	}
	return output
}