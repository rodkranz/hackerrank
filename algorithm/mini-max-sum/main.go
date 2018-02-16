package main

import (
	"fmt"
	"log"
)

func main() {
	nuns, err := readNumbers()
	if err != nil {
		log.Fatal(err)
	}
	
	min, max := calc(nuns)
	fmt.Printf("%v %v\n", min, max)
}

func calc(nuns []int64) (int64, int64) {
	var max int64 = 0
	var min int64 = 0
	
	for i := range nuns {
		var calc int64 = 0
		for n := 0; n < len(nuns); n++ {
			if n == i {
				continue
			}
			calc += nuns[n]
		}
		
		if max == 0 || calc < min {
			min = calc
		}
		if min == 0 || calc > max {
			max = calc
		}
	}
	
	return min, max
}

const (
	ConstraintsNumMin int64 = 1
	ConstraintsNumMax int64 = 1000000000
)

func readNumbers() ([]int64, error) {
	numbers := make([]int64, 5)
	for i := range numbers {
		_, err := fmt.Scanf("%d", &numbers[i])
		if err != nil {
			return nil, err
		}
		
		switch {
		case numbers[i] > ConstraintsNumMax:
			return nil, fmt.Errorf("the maximum value is %d", ConstraintsNumMax)
		case numbers[i] < ConstraintsNumMin:
			return nil, fmt.Errorf("the minimum value is %d", ConstraintsNumMin)
		}
	}
	return numbers, nil
}
