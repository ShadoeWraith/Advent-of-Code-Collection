package solutions

import (
	"strconv"
	"strings"
)

func PartOne(records []string) (int, error) {
	res := 0

	for _, record := range records {
		hyphenIndex := strings.Index(record, "-")
		r1 := record[:hyphenIndex]
		r2 := record[hyphenIndex+1:]

		start, _ := strconv.Atoi(r1)
		end, _ := strconv.Atoi(r2)

		for val := start; val <= end; val++ {
			s := strconv.Itoa(val)

			midFirstHalf := s[:len(s)/2]
			midSecondHalf := s[len(s)/2:]

			if midFirstHalf == midSecondHalf {
				res += val
			}
		}

	}

	return res, nil
}
