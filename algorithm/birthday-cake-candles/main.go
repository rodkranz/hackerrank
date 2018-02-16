package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	n, err := readNumber()
	if err != nil {
		log.Fatal(err)
	}

	ns, err := readNumbers(n)
	if err != nil {
		log.Fatal(err)
	}

	result := birthdayCakeCandles(ns)
	fmt.Println(result)
}

// birthdayCakeCandles execute birthdayCakeCandles of exercise.
func birthdayCakeCandles(ns []int) int {
	sort.Ints(ns)
	tallest := 0
	total := 0
	for i := len(ns) - 1; i >= 0; i-- {
		if tallest == 0 {
			tallest = ns[i]
		}

		if ns[i] != tallest {
			break
		}

		total += 1
	}
	return total
}

// Constraints
const (
	// Number
	ConstNumberMin = 1
	ConstNumberMax = 100000
	// Numbers
	ConstNumbersMin = 1
	ConstNumbersMax = 10000000
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
func readNumbers(n int) ([]int, error) {
	numbers := make([]int, n)
	for i := range numbers {
		_, err := fmt.Scanf("%d", &numbers[i])
		if err != nil {
			return nil, err
		}

		switch {
		case numbers[i] > ConstNumbersMax:
			return nil, fmt.Errorf("the maximum value is %d", ConstNumbersMax)
		case numbers[i] < ConstNumbersMin:
			return nil, fmt.Errorf("the minimum value is %d", ConstNumbersMin)
		}
	}
	return numbers, nil
}
