package solutions

import (
	"strconv"
	"strings"
)

func PartTwo(records []string) (int, error) {

	var res = 0

	for _, record := range records {
		hyphenIndex := strings.Index(record, "-")

		r1 := record[:hyphenIndex]
		r2 := record[hyphenIndex+1:]

		start, _ := strconv.Atoi(r1)
		end, _ := strconv.Atoi(r2)

		for val := start; val <= end; val++ {
			s := strconv.Itoa(val)

			if isInvalidID(s) {
				res += val
			}
		}
	}

	return res, nil
}

func isInvalidID(s string) bool {
	sLen := len(s)

	for i := 1; i <= sLen/2; i++ {
		if sLen%i == 0 {
			pattern := s[:i]
			isRepeat := true

			for j := i; j < sLen; j += i {
				segment := s[j : j+i]

				if segment != pattern {
					isRepeat = false
					break
				}
			}

			if isRepeat {
				return true
			}
		}
	}

	return false
}
