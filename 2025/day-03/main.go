package main

import (
	"AoC/2025/day-03/solutions"
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error reading from file, %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var records []string

	for scanner.Scan() {
		vals := scanner.Text()

		records = append(records, vals)
	}

	part1Res, _ := solutions.PartOne(records)
	part2Res, _ := solutions.PartTwo(records)

	fmt.Printf("Total Output for Part 1: %d\n", part1Res)
	fmt.Printf("Total Output for Part 2: %d\n", part2Res)
}
