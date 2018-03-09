package main

import (
	"fmt"
	"strings"
	"sort"
)

type words []string

func (p words) Shift() (string, words) {
	return p[0], p[1:]
}
func NewWord(input string) words {
	word := stringToArray(input)
	sort.Strings(word)
	return words(word)
}

func main() {
	l1 := readLine()
	l2 := readLine()
	
	result := anagram(l1, l2)
	fmt.Println(result)
}

func stringToArray(in string) []string {
	output := make([]string, len(in))
	for i := range strings.ToLower(in) {
		output[i] = string(in[i])
	}
	return output
}

func anagram(inputA, inputB string) int {
	ac := NewWord(inputA)
	bc := NewWord(inputB)
	
	var words words
	for len(ac) > 0 || len(bc) > 0 {
		if len(ac) == 0 {
			var x string
			x, bc = bc[0], bc[1:]
			words = append(words, x)
			continue
		}
		if len(bc) == 0 {
			var x string
			x, ac = ac[0], ac[1:]
			words = append(words, x)
			continue
		}
		var x string
		if ac[0] < bc[0] {
			x, ac = ac[0], ac[1:]
			words = append(words, x)
		} else if ac[0] > bc[0] {
			x, bc = bc[0], bc[1:]
			words = append(words, x)
		} else {
			ac = ac[1:]
			bc = bc[1:]
		}
	}
	return len(words)
}

func readLine() (line string) {
	fmt.Scan(&line)
	return line
}
