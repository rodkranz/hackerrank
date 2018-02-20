package main

import (
	"fmt"
	"log"
)

// Constraints
const (
	// Number
	constNumMin = 1
	constNumMax = 60
	// Numbers
	constNumsMin = 0
	constNumsMax = 100
)

// readNumber read os.Stdin and return the single number
func readNumber() (int) {
	var size int
	_, err := fmt.Scanf("%d", &size)
	
	switch {
	case err != nil:
		log.Fatal(err)
	case size > constNumMax:
		log.Fatalf("the maximum value is %d", constNumMax)
	case size < constNumMin:
		log.Fatalf("the minimum value is %d", constNumMin)
	}
	
	return size
}

func readNumbers(n int) ([]int) {
	numbers := make([]int, n)
	for i := range numbers {
		_, err := fmt.Scanf("%d", &numbers[i])
		if err != nil {
			log.Fatal(err)
		}
		
		switch {
		case numbers[i] > constNumsMax:
			log.Fatalf("the maximum value is %d", constNumsMax)
		case numbers[i] < constNumsMin:
			log.Fatalf("the minimum value is %d", constNumsMin)
		}
	}
	return numbers
}

func solve(ns []int) []int {
	n := make([]int, len(ns))
	for i := range ns {
		newGrade := ((ns[i] / 5) + 1) * 5
		
		if newGrade < 40 {
			n[i] = ns[i]
			continue
		}
		
		if (newGrade - ns[i]) < 3 {
			n[i] = newGrade
		} else {
			n[i] = ns[i]
		}
	}
	return n
}

func main() {
	n := readNumber()
	ns := readNumbers(n)
	
	result := solve(ns)
	for i := range result {
		fmt.Println(result[i])
	}
}
