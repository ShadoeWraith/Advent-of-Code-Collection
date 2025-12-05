package main

import (
	"fmt"
	"strings"

	"AoC/2025/day-02/solutions"
	"AoC/utils"
)

func main() {
	records := utils.ReadFromFile()

	for i := range records {
		vals := strings.Split(records[i], ",")
		records = append(records, vals...)
	}

	part1Res, _ := solutions.PartOne(records)
	part2Res, _ := solutions.PartTwo(records)

	fmt.Printf("Part1: Invalid IDs value added up: %d\n", part1Res)
	fmt.Printf("Part2: Invalid IDs value added up: %d\n", part2Res)
}
