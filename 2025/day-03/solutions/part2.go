package solutions

import "strconv"

func PartTwo(records []string) (int, error) {
	res := 0

	for _, record := range records {
		joltage := [2]int{}
		highestDigitIndex := 0

		for i := 0; i < len(record)-1; i++ {
			digit := int(record[i] - '0')

			if digit > joltage[0] {
				joltage[0] = digit
				highestDigitIndex = i
			}
		}

		subSlice := record[highestDigitIndex+1:]

		for i := range subSlice {
			digit := int(subSlice[i] - '0')

			if digit > joltage[1] {
				joltage[1] = digit
			}
		}

		resString := ""

		for i := range joltage {
			resString += strconv.Itoa(joltage[i])
		}

		resInt, _ := strconv.Atoi(resString)

		res += resInt
	}

	return res, nil
}
