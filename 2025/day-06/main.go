package main

import (
	"AoC/2025/day-06/solutions"
	"AoC/utils"
	"fmt"
	"strconv"
	"strings"
)

func isNumericLine(record string) bool {
	fields := strings.Fields(record)
	if len(fields) == 0 {
		return false
	}

	for _, field := range fields {
		if _, err := strconv.Atoi(field); err != nil {
			return false
		}
	}
	return true
}

func main() {
	records := utils.ReadFromFile()
	var numericData solutions.Calculation

	for _, record := range records {

		if isNumericLine(record) {

			numberStrings := strings.Split(record, " ")
			var currentLineNumbers []int

			for _, s := range numberStrings {
				s = strings.TrimSpace(s)
				if s == "" {
					continue
				}

				num, _ := strconv.Atoi(s)
				currentLineNumbers = append(currentLineNumbers, num)
			}

			if len(currentLineNumbers) > 0 {
				numericData.Numbers = append(numericData.Numbers, currentLineNumbers)
			}

		} else if strings.ContainsAny(record, "+*") {
			operators := strings.Split(record, " ")

			for _, op := range operators {
				op = strings.TrimSpace(op)

				if op == "+" || op == "*" {
					numericData.Operations = append(numericData.Operations, op)
				}
			}
		}
	}

	fmt.Printf("Solution to part 1: %v\n", solutions.PartOne(numericData))
}
