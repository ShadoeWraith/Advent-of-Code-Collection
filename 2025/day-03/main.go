package main

import (
	"AoC/2025/day-03/solutions"
	"AoC/utils"
	"fmt"
)

func main() {
	records := utils.ReadFromFile()

	part1Res, _ := solutions.PartOne(records)
	part2Res, _ := solutions.PartTwo(records)

	fmt.Printf("Total Output for Part 1: %d\n", part1Res)
	fmt.Printf("Total Output for Part 2: %d\n", part2Res)
}
