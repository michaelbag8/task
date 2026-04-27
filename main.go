package main

import (
	"fmt"
	"os"
	"strings"
)

func CountLinesInText(input string) int {
	sl := strings.Split(input, "\n")
	if input == "" {
		return 0
	}
	return len(sl)
}

func ReplaceWordInFile(file, newWord, oldWord string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}
	content := string(data)

	searchWord := strings.Fields(content)
	for i := 0; i < len(searchWord); i++ {
		if searchWord[i] != oldWord {
			continue
		}
		if searchWord[i] == oldWord {
			searchWord[i] = newWord
		}
	}
	return strings.Join(searchWord, " ")
}

func CountWords(str string) int {
	words := strings.Fields(str)
	return len(words)
}

func ReplaceAllInText(text, oldW, newW string) string {
	return strings.ReplaceAll(text, oldW, newW)
}

func PrintLineSegments(str string) error {
	words := strings.Split(str, "\\n")
	for i, word := range words {
		fmt.Printf("Line %d : %s\n", i, word)
	}
	return nil
}

func ValidateASCIIInput(str string) {
	for _, ch := range str {
		if ch < ' ' || ch > '~' {
			fmt.Fprintf(os.Stderr, "Invalid character found: %c\n", ch)
			return
		}
	}
	fmt.Println("Valid input")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Input the current file format")
		os.Exit(1)
	}

	input := os.Args[1]

	ValidateASCIIInput(input)

	lineCount := CountLinesInText(input)
	fmt.Printf("Line Count: %d\n", lineCount)

	wordCount := CountWords(input)
	fmt.Printf("Word Count: %d\n", wordCount)

	err := PrintLineSegments(input)
	if err != nil {
		fmt.Println("Error printing line segments")
		os.Exit(1)
	}

	if len(os.Args) >= 4 {
		oldWord := os.Args[2]
		newWord := os.Args[3]

		replaced := ReplaceAllInText(input, oldWord, newWord)
		fmt.Printf("Replaced Text: %s\n", replaced)
	}

	// ReplaceWordInFile — replace a word in a file
	if len(os.Args) == 4 {
		file := os.Args[1]
		oldWord := os.Args[2]
		newWord := os.Args[3]

		updatedContent := ReplaceWordInFile(file, oldWord, newWord)
		fmt.Printf("Updated File Content:\n%s\n", updatedContent)
	}
}
