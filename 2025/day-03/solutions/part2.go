package solutions

import (
	"strconv"
)

// Used AI to help find solve this.
// Wasn't familiar with the "greedy algorithm" prior to this.

func PartTwo(records []string) (int, error) {
	totalJoltage := 0

	for _, record := range records {

		joltageString := ""

		const totalDigitsToSelect int = 12
		totalDigitsToOmit := len(record) - totalDigitsToSelect

		currentPos := 0

		omissionsRemaining := totalDigitsToOmit

		for i := range totalDigitsToSelect {
			digitsToComplete := totalDigitsToSelect - i

			searchLimit := len(record) - digitsToComplete

			limitIndex := currentPos + omissionsRemaining
			if searchLimit < limitIndex {
				limitIndex = searchLimit
			}

			bestDigit := 0
			bestIndex := 0

			for j := currentPos; j <= limitIndex; j++ {
				digit := int(record[j] - '0')

				if digit > bestDigit {
					bestDigit = digit
					bestIndex = j
				}
			}

			joltageString += strconv.Itoa(bestDigit)

			omitted := bestIndex - currentPos
			omissionsRemaining -= omitted

			currentPos = bestIndex + 1
		}

		resInt, _ := strconv.Atoi(joltageString)

		totalJoltage += resInt
	}

	return totalJoltage, nil
}
