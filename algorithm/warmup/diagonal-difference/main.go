package main

import (
	"fmt"
	"log"
	"math"
)

const (
	CONSTRAINTS_MIN = -100
	CONSTRAINTS_MAX = 100
)

var (
	constraintMinError = fmt.Errorf("the min num is %d", CONSTRAINTS_MIN)
	constraintMaxError = fmt.Errorf("the max num is %d", CONSTRAINTS_MAX)
)

func main() {
	n, err := readSize()
	if err != nil {
		log.Fatal(err)
	}

	matrix := make(map[int][]int, n)

	for i := 0; i < n; i++ {
		ns, err := readLines(n)
		if err != nil {
			log.Fatal(err)
		}

		matrix[i] = ns
	}

	A1 := 0
	A2 := 0

	in := len(matrix) - 1
	for i := 0; i < len(matrix); i++ {
		for n := 0; n < len(matrix[i]); n++ {
			if in == n {
				A2 += matrix[i][n]
			}
			if i == n {
				A1 += matrix[i][n]
			}
		}
		in--
	}

	fmt.Println(math.Abs(float64(A2 - A1)))
}

func readSize() (int, error) {
	var size int
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		return -1, err
	}
	return size, nil
}

func readLines(N int) ([]int, error) {
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		_, err := fmt.Scanf("%d", &arr[i])
		if err != nil {
			return nil, err
		}

		switch {
		case arr[i] < CONSTRAINTS_MIN:
			return nil, constraintMinError
		case arr[i] > CONSTRAINTS_MAX:
			return nil, constraintMaxError
		}
	}
	return arr, nil
}
