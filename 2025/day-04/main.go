package main

import (
	"fmt"

	"AoC/2025/day-04/solutions"
	"AoC/utils"
)

func main() {
	records := utils.ReadFromFile()

	part1Res, _ := solutions.PartOne(records)
	part2Res, _ := solutions.PartTwo(records)

	fmt.Printf("Part1: Total of Valid Rolls: %d\n", part1Res)
	fmt.Printf("Part2: Invalid IDs value added up: %d\n", part2Res)
}
