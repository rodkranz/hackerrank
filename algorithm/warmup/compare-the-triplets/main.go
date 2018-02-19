package main

import (
	"fmt"
	"log"
)

func main() {
	points := 3
	
	A, err := readLines(points)
	if err != nil {
		log.Fatal(err)
	}
	
	B, err := readLines(points)
	if err != nil {
		log.Fatal(err)
	}
	
	result := check(A, B)
	for i := range result {
		fmt.Printf("%d ", result[i])
	}
}

func check(A, B []int) []int {
	result := []int{0, 0}
	
	for i := range A {
		switch {
		case A[i] > B[i]:
			result[0]++
		case A[i] < B[i]:
			result[1]++
		}
	}
	return result
}

func readLines(N int) ([]int, error) {
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		_, err := fmt.Scanf("%d", &arr[i])
		if err != nil {
			return nil, err
		}
	}
	return arr, nil
}
