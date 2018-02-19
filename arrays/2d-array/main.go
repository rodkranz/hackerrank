package main

import (
	"fmt"
	"log"
)

func main() {
	size, err := readSize()
	if err != nil {
		log.Fatal(err)
	}
	
	numbers, err := readNumbers(size)
	if err != nil {
		log.Fatal(err)
	}
	
	printArray(numbers)
}

func printArray(arr []int) {
	for i := len(arr)-1; i >= 0; i-- {
		fmt.Printf("%d", arr[i])
		
		if i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func readSize() (int, error) {
	var length int
	
	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		return -1, err
	}
	
	return length, nil
}

func readNumbers(size int) ([]int, error) {
	numbers := make([]int, size)
	for i := range numbers {
		_, err := fmt.Scanf("%d", &numbers[i])
		if err != nil {
			return nil, err
		}
	}
	return numbers, nil
}
