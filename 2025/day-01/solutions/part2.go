package solutions

import (
	"strconv"
)

func PartTwo(records []string) (int, error) {
	res := 0
	cur := 50
	prev := 0

	for i := range records {
		line := records[i]
		val, _ := strconv.Atoi(line[1:])

		res += val / 100

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

		switch {
		case cur == 0:
			res++
		case cur >= prev && line[0] == 'L' && prev != 0:
			res++
		case cur <= prev && line[0] == 'R' && prev != 0:
			res++
		}
		prev = cur
	}

	return res, nil
}
