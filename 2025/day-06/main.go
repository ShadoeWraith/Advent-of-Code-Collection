package main

import (
	"AoC/2025/day-06/solutions"
	"AoC/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	records := utils.ReadFromFile()
	var numericData solutions.Calculation

	for _, record := range records {
		if strings.ContainsAny(record, "+*") {
			operators := strings.SplitSeq(record, " ")

			for op := range operators {
				numericData.Operations = append(numericData.Operations, op)
			}
			continue
		}

		numberStrings := strings.Split(record, " ")
		var currentLineNumbers []int

		for _, s := range numberStrings {
			num, _ := strconv.Atoi(s)
			currentLineNumbers = append(currentLineNumbers, num)
		}

		numericData.Numbers = append(numericData.Numbers, currentLineNumbers)
	}

	fmt.Printf("Solution to part 1: %v\n", solutions.PartOne(numericData))
	fmt.Printf("Solution to part 2: %v\n", solutions.PartTwo(numericData))
}
