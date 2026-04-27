package main

import (
	"fmt"
	"os"
	"strings"
)

// loadBanner reads the font file and returns a map of rune -> 8 lines of ASCII art
func loadBanner(filename string) map[rune][]string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading fontfile: %v\n", err)
		os.Exit(1)
	}

	content := string(data)
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.TrimLeft(content, "\n")

	rawFile := strings.Split(content, "\n\n")

	bannerMap := make(map[rune][]string)
	charCode := rune(' ') // ASCII art files start at space (32)

	for _, raw := range rawFile {
		lines := strings.Split(raw, "\n")
		if len(lines) < 8 {
			continue
		}
		bannerMap[charCode] = lines[:8]
		charCode++
	}

	return bannerMap
}

// renderBanner uses the banner map to print the ASCII art for the input string
func renderBanner(input string, bannerMap map[rune][]string) {
	segments := strings.Split(input, "\\n")

	for i, segment := range segments {
		if segment == "" {
			if i < len(segments)-1 {
				fmt.Println()
			}
			continue
		}

		// Validate all characters before rendering
		for _, ch := range segment {
			if _, ok := bannerMap[ch]; !ok {
				fmt.Fprintf(os.Stderr, "Non Printable ASCII Character %q\n", ch)
				os.Exit(1)
			}
		}

		// Render row by row
		for row := 0; row < 8; row++ {
			var result strings.Builder
			for _, ch := range segment {
				result.WriteString(bannerMap[ch][row])
			}
			fmt.Println(result.String())
		}
	}
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Fprintf(os.Stderr, "Usage: go run . text or go run . text banner\n")
		os.Exit(1)
	}

	fontFile := "standard.txt"
	input := os.Args[1]

	if len(os.Args) == 3 {
		fontFile = os.Args[2] + ".txt"
	}

	bannerMap := loadBanner(fontFile)
	renderBanner(input, bannerMap)
}