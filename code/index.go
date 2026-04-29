package main

import (
	"fmt"
	"os"
	"strings"
)

// ListBanners scans the given directory and returns the names
// of all valid banner files (files ending in ".txt"),
// stripped of their extension. e.g ["standard", "shadow", "thinkertoy"]
func ListBanners(dir string) ([]string, error) {
	data, err := os.ReadDir(dir)
	if err !=nil{
		fmt.Fprintf(os.Stderr, "Error geting directory %v", err)
	}

	var result []string

	for _, entries := range data{
		name := entries.Name()
		if !entries.IsDir() && strings.HasSuffix(name, ".txt"){
			result = append(result, strings.TrimSuffix(name, ".txt"))
		}
	}
	return  result, nil

}
// BannerSupports returns true if all characters in the given string
// are present as keys in the bannerMap, false otherwise.
// It should ignore the literal "\n" sequence when checking.
func BannerSupports(input string, bannerMap map[rune][]string) bool {
	input = strings.ReplaceAll(input, `\n`, "")
	for _, v := range input {
		if _, ok := bannerMap[v]; !ok {
			return false
		}
	}
	return true

}

// ValidateInput checks if all characters in the input string
// (excluding literal "\n") are printable ASCII (32–126).
// Returns an error if any invalid character is found.
func ValidateInput(input string) error {
	for _, ch := range input {
		if ch < 32 || ch > 126 {
			return fmt.Errorf("Error not a valid character %c", ch)
		}
	}
	return nil
}

//SplitInput slip user input 
//including literal "\n" by turning into a real newline 
func SplitInput(str string) []string {
	var output []string
	segments := strings.Split(str, "\\n")
	for _, segment := range segments {
		if segment == "" {
			continue
		}
		output = append(output, segment)
	}
	return output
}
