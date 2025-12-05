package main

import (
	"AoC/2025/day-05/solutions"
	"AoC/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	records := utils.ReadFromFile()

	var inputIds []int
	var minIds []int
	var maxIds []int
	min := 0
	max := 0

	for _, record := range records {
		hyphenIndex := strings.Index(record, "-")

		if hyphenIndex != -1 {
			min, _ = strconv.Atoi(record[:hyphenIndex])
			max, _ = strconv.Atoi(record[hyphenIndex+1:])

			minIds = append(minIds, min)
			maxIds = append(maxIds, max)

		} else {
			input, _ := strconv.Atoi(record)

			if input != 0 {
				inputIds = append(inputIds, input)
			}
		}
	}

	partOneRes, _ := solutions.PartOne(minIds, maxIds, inputIds)
	partTwoRes, _ := solutions.PartTwo(minIds, maxIds)

	fmt.Printf("The Fresh Ingredient Count For Part 1 is: %d\n", partOneRes)
	fmt.Printf("The Fresh Ingredient Count For Part 2 is: %d", partTwoRes)
}
