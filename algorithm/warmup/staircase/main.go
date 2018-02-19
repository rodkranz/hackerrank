package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	size, err := readSize()
	if err != nil {
		log.Fatal(err)
	}
	
	for n := 0; n < size; n++ {
		for i := size - n; i > 1; i-- {
			fmt.Print(" ")
		}
		
		fmt.Println(strings.Repeat("#", n+1))
	}
}

func readSize() (int, error) {
	var length int
	
	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		return -1, err
	}
	
	return length, nil
}
