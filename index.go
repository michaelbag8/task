package main

import (
	"fmt"
	"strings"
)

func printBanner(str string) (blocks [][]string) {
	words := strings.Split(str, "\n\n")
	for _, ch := range words {
		lines := strings.Split(ch, "\n")
		blocks = append(blocks, lines)
	}
	return blocks
}

func BuildIndexMap() map[rune]int {
	blocks := make(map[rune]int)

	for ch := ' '; ch <= '~'; ch++ {
		index := int(ch - ' ')
		blocks[ch] = index
	}
	return blocks

}

func WordFrequency(text string) map[string]int { // check on this
	count := make(map[string]int)

	words := strings.Fields(text)

	for index, word := range words {
		count[word] = index
	}
	return count
}

func BannerToMap(blocks [][]string) map[rune][]string {
	block := make(map[rune][]string)

	for index, char := range blocks {
		ch := rune(index + int(' '))
		block[ch] = char
	}
	return block
}

func SlicetoGrid(str string, cols int) (blocks [][]string) {
	words := strings.Fields(str)

	for row := 0; row < len(words); row += cols {
		end := row + cols
		if end > len(words) {
			end = len(words)
		}
		blocks = append(blocks, words[row:end])
	}
	return blocks
}

func main() {
	words := "mango banana berry apple guava banana berry kiwi apple guava banana berry"
	text := "go is fun go is powerful"
	fmt.Println(WordFrequency(words))
	fmt.Println(WordFrequency(text))

	//PrintBanner
	banner := "line1\nline2\nline3\n\nline4\nline5\nline6"
	fmt.Println(printBanner(banner))

	//BuildIndexMap prints all the printable ASCII character from '' to '~'
	m := BuildIndexMap()
	fmt.Println("Index of 'A':", m['A'])
	fmt.Println("Index of 'a':", m['a'])
	fmt.Println("Index of '0':", m['0'])
	fmt.Println("Index of ' ':", m[' '])

	for ch := 'A'; ch <= 'D'; ch++ {
		fmt.Printf("%c -> %d\n", ch, m[ch])
	}

	board := "X _ O _ X _ O _ X"
	fmt.Println(StringToGrid(board))
	wrd := "one two three four five six seven eight nine ten eleven twelve"
	fmt.Println(SlicetoGrid(wrd, 3))

}
