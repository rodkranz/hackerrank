package main

import (
	"fmt"
	"log"
	"unicode"
)

func main() {
	line, err := readLine()
	if err != nil {
		log.Fatal(err)
	}
	
	result := camelCase(line)
	fmt.Println(result)
}

func camelCase(input string) int {
	var result int
	
	if len(input) > 0 {
		result = 1
	}
	
	for _, r := range input {
		if unicode.IsUpper(r) {
			result += 1
		}
	}
	
	return result
}

func readLine() (line string, err error) {
	_, err = fmt.Scan(&line)
	return line, err
}
