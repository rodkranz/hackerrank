package main

import (
	"fmt"
	"log"
)

func main() {
	line, err := readLine()
	if err != nil {
		log.Fatal(err)
	}
	
	result := reduce(line)
	fmt.Println(result)
}

func reduce(input string) (string) {
	var out string
	for i := 0; i < len(input); i++ {
		if len(out) > 0 && out[len(out)-1] == input[i] {
			out = out[0 : len(out)-1]
		} else {
			out += string(input[i])
		}
	}
	
	if len(out) == 0 {
		out = "Empty String"
	}
	
	return out
}

func readLine() (line string, err error) {
	_, err = fmt.Scan(&line)
	return line, err
}
