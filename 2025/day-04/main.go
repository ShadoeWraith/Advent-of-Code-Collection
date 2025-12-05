package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"AoC/2025/day-04/solutions"
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

	part1Res, _ := solutions.PartOne(records)
	part2Res, _ := solutions.PartTwo(records)

	fmt.Printf("Part1: Total of Valid Rolls: %d\n", part1Res)
	fmt.Printf("Part2: Invalid IDs value added up: %d\n", part2Res)
}
