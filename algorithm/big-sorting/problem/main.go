package main

import (
	"fmt"
	"log"
)

func main() {
	n, err := readSize()
	checkErr(err)
	ns, err := readLines(n)
	checkErr(err)

	fmt.Println("Unsorted:", ns)
	bubbleSort(ns)
	fmt.Println("Sorted:", ns)
}

func bubbleSort(numbers []uintptr) {
	N := len(numbers)
	var i int
	for i = 0; i < N; i++ {
		if !sweep(numbers, i) {
			return
		}
	}
}

func sweep(numbers []uintptr, prevPasses int) bool {
	N := len(numbers)
	firstIndex, secondIndex, didSwap := 0, 1, false

	for secondIndex < (N - prevPasses) {
		var firstNumber = numbers[firstIndex]
		var secondNumber = numbers[secondIndex]

		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
			didSwap = true
		}

		firstIndex++
		secondIndex++
	}

	return didSwap
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readSize() (int, error) {
	var size int
	_, err := fmt.Scan(&size)
	if err != nil {
		return -1, err
	}
	return size, nil
}

func readLines(n int) ([]int64, error) {
	arr := make([]int64, n)
	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&arr[i]); err != nil {
			return nil, err
		}
	}
	return arr, nil
}
