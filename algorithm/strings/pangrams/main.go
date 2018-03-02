package main

import (
	"fmt"
	"log"
	"strings"
	"io"
)

var words map[int32]bool

func init() {
	words = make(map[int32]bool, 25)
	for i := 97; i < 122; i++ {
		words[int32(i)] = false
	}
}

func main() {
	line, err := readLine()
	if err != nil {
		log.Fatal(err)
	}
	
	line = strings.ToLower(line)
	if pangram(line) {
		fmt.Println("pangram")
	} else {
		fmt.Println("not pangram")
	}
}

func pangram(input string) (bool) {
	for _, v := range input {
		words[v] = true
	}
	
	for _, b := range words {
		if !b {
			return false
		}
	}
	
	return true
}

func readLine() (line string, err error) {
	final := ""
	for {
		_, err = fmt.Scan(&line)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		final += line
		line = ""
	}
	return final, err
}
