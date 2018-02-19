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
	
	out := fraction(float64(size), numbers)
	for i := range out {
		fmt.Printf("%.6f\n", out[i])
	}
}

func fraction(size float64, n []int) []float64 {
	var positive float64 = 0
	var negative float64 = 0
	var zeroes float64 = 0
	
	for i := range n {
		switch {
		case n[i] > 0:
			positive += 1
		case n[i] < 0:
			negative += 1
		default:
			zeroes += 1
		}
	}
	
	return []float64{
		positive / size,
		negative / size,
		zeroes / size,
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
