package main

import (
	"fmt"
	"log"
)

func main() {
	n, err := readSize()
	if err != nil {
		log.Fatal(err)
	}

	ns, err := readLines(n)
	if err != nil {
		log.Fatal(err)
	}

	total := sumElements(ns)
	fmt.Printf("%d", total)
}

func sumElements(elms []int) int {
	var total int
	for i := range elms {
		total += elms[i]
	}
	return total
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
	}
	return arr, nil
}
