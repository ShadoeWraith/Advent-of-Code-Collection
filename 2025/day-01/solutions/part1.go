package solutions

import (
	"strconv"
)

// Part One =====================================================

func PartOne(records []string) (int, error) {
	res := 0
	cur := 50

	for i := range records {
		line := records[i]
		val, _ := strconv.Atoi(line[1:])

		val %= 100
		if line[0] == 'L' {
			val = -val
		}

		cur += val

		if cur > 99 {
			cur -= 100
		} else if cur < 0 {
			cur += 100
		}
		if cur == 0 {
			res++
		}
	}

	return res, nil
}
