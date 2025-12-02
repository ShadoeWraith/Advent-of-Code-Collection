package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filepath := "input.txt"

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var records []string

	for scanner.Scan() {
		vals := scanner.Text()

		newVal := strings.Split(vals, ",")

		records = append(records, newVal...)
	}

	res, _ := partOne(records)

	fmt.Printf("Part1: Invalid IDs value added up: %d\n", res)
}

func partOne(records []string) (int, error) {
	res := 0

	for _, record := range records {
		hyphenIndex := strings.Index(record, "-")
		r1 := record[:hyphenIndex]
		r2 := record[hyphenIndex+1:]

		start, _ := strconv.Atoi(r1)
		end, _ := strconv.Atoi(r2)

		for val := start; val <= end; val++ {
			s := strconv.Itoa(val)

			midFirstHalf := s[:len(s)/2]
			midSecondHalf := s[len(s)/2:]

			if midFirstHalf == midSecondHalf {
				res += val
			}
		}

	}

	return res, nil
}
