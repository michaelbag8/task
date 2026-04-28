package main

import (
	"fmt"
	"strings"
)

func firstAndLast(str string) (string, string) {
	words := strings.Fields(str)

	if len(words) == 0 {
		return "", ""
	}
	return words[0], words[len(words)-1]
}

func wordAtIndex(str string, index int) (string, bool) {
	words := strings.Fields(str)
	if index < 0 || index >= len(words) {
		return "", false
	}

	return words[index], true
}
func safeGetRune(s string, i int) rune {
	runes := []rune(s)
	if i < 0 || i >= len(runes) {
		return 0
	}
	return runes[i]
}


func prevWords(sentence string, i, n int) []string {
	words := strings.Fields(sentence)
	start := i - n

	if start < 0 {
		start = 0
	}
	if i <= 0 || i >= len(words) {
		return []string{}
	}
	word := words[start:i]
	return word

}

// returns the n words after position i
func nextWords(sentence string, i, n int) []string {
	words := strings.Fields(sentence)

	if i < 0 || i >= len(words) {
		return []string{}
	}

	start := i
	end := i + 1 + n
	if end > len(words) {
		end = len(words)
	}

	return words[start:end]
}

func nWordsBefore(sentence string, i, n int) string {
	words := strings.Fields(sentence)
	if i <= 0 || i > len(words) {
		return " "
	}

	start := i - n
	if start < 0 {
		start = 0
	}

	return strings.Join(words[start:i], " ")

}


func main() {

	//firstAndLast words
	f, l := firstAndLast("hello")
	fmt.Println(f, l)

	//wordAtIndex
	s := "Go is fun"
	word, ok := wordAtIndex(s, 2)
	if ok {
		fmt.Printf("Found: %s\n", word)

	} else {
		fmt.Println("index out of range")
	}

	//safeGetRune
	g := "café"
	fmt.Println(safeGetRune(g, 0))
	r := "hello"
	fmt.Println(safeGetRune(r, 2))

	sentence := "we offer them bread but they were counting the gold were we trade"
	fmt.Println(prevWords(sentence, 5, 2))

	z := "the quick brown fox"
	fmt.Println(nextWords(z, 0, 1))

	fmt.Println(nWordsBefore("one two three", 1, 5))


}
