package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string,error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading fontfile: %v\n", err)
		return nil, err
	}

	if len(data) == 0{
		return nil, fmt.Errorf("banner can not be empty")
	}
	
	content := string(data)
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.TrimLeft(content, "\n")

	rawFile := strings.Split(content, "\n\n")

	bannerMap := make(map[rune][]string)
	charCode := rune(' ')

	for _, raw := range rawFile {
		lines := strings.Split(raw, "\n")
		if len(lines) < 8 {
			return nil, fmt.Errorf("invalid character")
		}
		bannerMap[charCode] = lines[:8]
		charCode++
	}

	return bannerMap, nil
}

func renderBanner(input string, bannerMap map[rune][]string) {
	segments := strings.Split(input, "\\n")

	for i, segment := range segments {
		if segment == "" {
			if i < len(segments)-1 {
				fmt.Println()
			}
			continue
		}

		for _, ch := range segment {
			if _, ok := bannerMap[ch]; !ok {
				fmt.Fprintf(os.Stderr, "Non Printable ASCII Character %q\n", ch)
				os.Exit(1)
			}
		}

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

	bannerMap, err := LoadBanner(fontFile)
	if err!=nil{
		fmt.Println("Error loading banner function")
		os.Exit(1)
	}
	renderBanner(input, bannerMap)
}
