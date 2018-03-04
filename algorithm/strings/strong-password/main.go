package main

import (
	"fmt"
	"log"
	"unicode"
)

func main() {
	num, err := readNumber()
	if err != nil {
		log.Fatal(err)
	}
	
	line, err := readLine()
	if err != nil {
		log.Fatal(err)
	}
	
	result := passwordStrong(num, line)
	fmt.Println(result)
}

func passwordStrong(_ int, input string) (int) {
	addCount := 0
	sum := map[string]int{
		"dc": 0,
		"lc": 0,
		"uc": 0,
		"sc": 0,
	}
	
	for _, c := range input {
		switch {
		case unicode.IsLower(c):
			sum["lc"]++
		case unicode.IsUpper(c):
			sum["uc"]++
		case unicode.IsNumber(c):
			sum["dc"]++
		default:
			sum["sc"]++
		}
	}
	
	for key := range sum {
		if sum[key] == 0 {
			addCount++
			sum[key]++
		}
	}
	
	total := 0
	for key := range sum {
		total += sum[key]
	}
	
	if total-6 < 0 {
		addCount += 6 - total
	}
	
	return addCount
}

func readNumber() (num int, err error) {
	_, err = fmt.Scan(&num)
	return num, err
}

func readLine() (line string, err error) {
	_, err = fmt.Scan(&line)
	if len(line) < 1 || len(line) > 100 {
		return "", fmt.Errorf("constraints is minimum 1 and maximum 100")
	}
	
	return line, err
}
