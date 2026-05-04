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

func findWord(sentence, target string) int {
    words := strings.Fields(sentence)
    for i, w := range words {
        if strings.EqualFold(w, target) {
            return i    
        }
    }
    return -1           
}

func findAllIndexes(sentence, target string) []int {
    words := strings.Fields(sentence)
    var result []int                      

    for i, w := range words {
        if strings.EqualFold(w, target) {
            result = append(result, i)    
        }
    }
    return result
}
func replaceAt(sentence, replacement string, i int) string {
    words := strings.Fields(sentence)
    if i < 0 || i >= len(words) {
        return sentence
    }
    words[i] = replacement
    return strings.Join(words, " ")
}



func replaceALl(sentence, replacement string, index int) string {
	words := strings.Fields(sentence)
	if index < 0 || index >= len(words) {
		index = len(words)
	}

	words[index] = replacement
	return strings.Join(words, " ")
}

func replaceFirst(sentence, replacement string) string {
	words := strings.Fields(sentence)
	if len(words) == 0 {
		return sentence
	}

	words[0] = replacement

	return strings.Join(words, " ")
}

func replaceWordBefore(sentence, target, replacement string) string {
	words := strings.Fields(sentence)
	for i := 0; i < len(words); i++ {
		if strings.EqualFold(words[i], target) && i > 0 {
			words[i-1] = replacement
			break
		}

	}

	return strings.Join(words, " ")
}

func replaceWordAfter(sentence, target, replacement string) string {
	words := strings.Fields(sentence)

	for i, word := range words {

		if word == target && i+1 < len(words) {
			words[i+1] = replacement
			break
		}
	}

	return strings.Join(words, " ")
}

func replaceLast(sentence, replacement string) string {
	words := strings.Fields(sentence)
	if len(words) == 0 {
		return sentence
	}

	words[len(words)-1] = replacement

	return strings.Join(words, " ")
}

func replaceWords(sentence, target, replacement string) string {
	words := strings.Fields(sentence)
	for i := 0; i < len(words); i++ {
		if i < 0 || i < len(words) {
			if strings.EqualFold(words[i], target) {
				words[i] = replacement
			}
		}

	}

	return strings.Join(words, " ")
}

func changeChar(s string, i int, r rune) string {
	runes := []rune(s)
	if i < 0 || i >= len(runes) {
		return s
	}

	runes[i] = r

	return string(runes)
}
func changeMultiplesOfFive(s string, r rune) string {
	runes := []rune(s)
	for i := range runes {
		if i%5 == 0 {
			runes[i] = r
		}

	}
	return string(runes)
}


func changeEvery5thWord(sentence string) string {
	words := strings.Fields(sentence)
	for i := range words {
		if i != 0 && i%5 == 0 {
			words[i] = strings.ToUpper(words[i])
		}
	}
	return strings.Join(words, " ")
}


func upperInsideQuotes(s string) string {
	words := strings.Fields(s)
	inQuote := false
	for i, w := range words {
		if w == "'" {
			inQuote = !inQuote
			continue

		}
		if inQuote {
			words[i] = strings.ToUpper(words[i])

		}
	}
	return strings.Join(words, " ")
}

func lowerOutsideQuotes(s string) string {
	words := strings.Fields(s)
	inQuote := false
	for i, w := range words {
		if w == "'" {
			inQuote = !inQuote
			continue

		}

		if !inQuote {
			words[i] = strings.ToLower(words[i])

		}

	}
	return strings.Join(words, " ")
}

func titleLowerCase(s string) string {
	words := strings.Fields(s)
	IsQuote := false
	for i, w := range words {
		if w == "'" {
			IsQuote = !IsQuote
			continue

		}
		if IsQuote {
			words[i] = strings.ToLower(words[i])
		} else {
			runes := []rune(words[i])
			if len(runes) > 0 {
				words[i] = strings.ToUpper(string(runes[0])) + strings.ToLower(string(runes[1:]))
			}
		}

	}
	return strings.Join(words, " ")
}
func replaceInsideQuotes(s, target, replacement string) string {
	words := strings.Fields(s)
	isQuote := false
	for i, w := range words {
		if w == "'" {
			isQuote = !isQuote
			continue
		}
		if isQuote && strings.EqualFold(w, target) {
			words[i] = replacement
		}

	}
	return strings.Join(words, " ")
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

	
    sw := "the quick brown fox"
    fmt.Println(findWord(sw, "brown")) 
    fmt.Println(findWord(sw, "BROWN")) 
    fmt.Println(findWord(sw, "cat"))  

	sd := "the cat sat on the mat near the cat"
    fmt.Println(findAllIndexes(sd, "the")) 
    fmt.Println(findAllIndexes(sd, "cat")) 
    fmt.Println(findAllIndexes(sd, "dog")) 
	

    sentences := "the quick brown fox"
    fmt.Println(replaceAt(sentences, "red", 2))  
    fmt.Println(replaceAt(sentences, "slow", 1)) 


	words := "This is a good way to start learning"
	fmt.Println(replaceALl(words, "earning", 7))
	fmt.Println(replaceFirst(words, "That"))
	fmt.Println(replaceLast(words, "playing"))

	fmt.Println(replaceWordBefore(words, "good", "great"))

	sz := "I love coding in Go every day"
	fmt.Println(replaceWordBefore(sz, "Go", "writing"))

	fmt.Println(replaceWordAfter(sz, "Go", "some"))

	w := "the cat and the dog and the bird"
	fmt.Println(replaceWords(w, "the", "a"))


	core := "hello"
	fmt.Println(changeChar(core, 1, 'a'))
	
	fmt.Println(changeMultiplesOfFive("abcdefghijunbvcdfreswaziokhg", 'T'))


	sen := "one two three four five six seven eight nine ten eleven twelve thirteen forteen"
	fmt.Println(changeEvery5thWord(sen))

	fmt.Println(upperInsideQuotes("Jesus deserve your ' praise all the time '"))
	fmt.Println(upperInsideQuotes("I saw ' hello world ' today"))
	fmt.Println(lowerOutsideQuotes(" THIS IS A PLACE TO ' LEARN CODING '"))
	fmt.Println(lowerOutsideQuotes("THE QUICK ' brown fox ' JUMPS"))
	
	fmt.Println(titleLowerCase("jesus deserve your ' PRAISE ALL THE TIME ' praise ye the lord"))



}
