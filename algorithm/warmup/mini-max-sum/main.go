package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	nuns, err := readNumbers()
	if err != nil {
		log.Fatal(err)
	}

	min, max := calc(nuns)
	fmt.Printf("%v %v\n", min, max)
}

func calc(nuns []int) (int, int) {
	sort.Ints(nuns)

	max := 0
	min := 0

	for i, n := 0, len(nuns)-1; i < len(nuns)-1; i++ {
		min += nuns[i]
		max += nuns[n]

		n -= 1
	}

	return min, max
}

const (
	constNumMin = 1
	constNumMax = 1000000000
)

func readNumbers() ([]int, error) {
	numbers := make([]int, 5)
	for i := range numbers {
		_, err := fmt.Scanf("%d", &numbers[i])
		if err != nil {
			return nil, err
		}

		switch {
		case numbers[i] > constNumMax:
			return nil, fmt.Errorf("the maximum value is %d", constNumMax)
		case numbers[i] < constNumMin:
			return nil, fmt.Errorf("the minimum value is %d", constNumMin)
		}
	}
	return numbers, nil
}
