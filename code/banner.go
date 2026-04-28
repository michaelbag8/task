package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(file string) (map[rune][]string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file")
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("file cannot be empty")
	}
	content := string(data)
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.TrimLeft(content, "\n")

	blocksMap := make(map[rune][]string)
	charCode := ' '

	rawFile := strings.Split(content, "\n\n")

	for _, raw := range rawFile {
		lines := strings.Split(raw, "\n")
		if len(lines) < 8 {
			return nil, fmt.Errorf("Invalid file")
		}
		blocksMap[charCode] = lines[:8]
		charCode++
	}

	return blocksMap, nil
}