package main

import (
	"fmt"
	"log"
	"unicode"
)

func main() {
	num, err := readNumber()
	if err != nil {
		log.Fatal(err)
	}
	
	input, err := readStringBytes(num)
	if err != nil {
		log.Fatal(err)
	}
	
	result := twoCharacters(input)
	fmt.Printf("%v\n", result)
	
}

func twoCharacters(input string) int {
	var result int
	
	for i := 0; i < len(input); i++ {
	
	}
	
	
	return result
}

// Constraints
const (
	// Number
	ConstNumberMin = 1
	ConstNumberMax = 1000
)

// readNumber read os.Stdin and return the single number
func readNumber() (int, error) {
	var size int
	_, err := fmt.Scanf("%d", &size)
	
	switch {
	case err != nil:
		return -1, err
	case size > ConstNumberMax:
		return -1, fmt.Errorf("the maximum value is %d", ConstNumberMax)
	case size < ConstNumberMin:
		return -1, fmt.Errorf("the minimum value is %d", ConstNumberMin)
	}
	
	return size, nil
}

// readNumbers read os.Stdin and return the list of numbers
func readStringBytes(n int) (string, error) {
	var input string
	
	if _, err := fmt.Scanf("%s", &input); err != nil {
		return "", err
	}
	
	for _, v := range input {
		if !unicode.IsLetter(v) {
			return "", fmt.Errorf("the allowed is only letters %s is not valid", string(v))
		}
	}
	
	return input, nil
}
