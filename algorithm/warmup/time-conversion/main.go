package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	line, err := readLine()
	if err != nil {
		log.Fatal(err)
	}

	result, err := convert(line)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func convert(input string) (string, error) {
	formatA := "03:04:05PM"
	formatB := "15:04:05"

	t, err := time.Parse(formatA, input)
	if err == nil {
		return t.Format(formatB), nil
	}

	t, err = time.Parse(formatB, input)
	if err == nil {
		return t.Format(formatA), nil
	}

	return "", fmt.Errorf("format unknown")
}

func readLine() (line string, err error) {
	_, err = fmt.Scanf("%s", &line)
	return line, err
}
